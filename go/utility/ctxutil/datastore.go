package ctxutil

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	readerDB     *sql.DB
	writerDB     *sql.DB
	readerDBStmt *sync.Map
	writerDBStmt *sync.Map
)

type DataStore struct {
	readerDB    *sql.DB
	writerDB    *sql.DB
	transaction *sql.Tx
}

func initDatastore() {
	var err error
	readerDSN := os.Getenv("DSN") //os.Getenv("MYSQL_READ_DSN")
	readerDB, err = sql.Open("mysql", readerDSN)
	if err != nil {
		panic(err)
	}
	readerDB.SetMaxOpenConns(10)
	readerDB.SetMaxIdleConns(10)
	// readerDB.SetConnMaxLifetime(time.Second * 1)

	writerDSN := os.Getenv("DSN") // os.Getenv("MYSQL_WRITE_DSN")
	writerDB, err = sql.Open("mysql", writerDSN)
	if err != nil {
		panic(err)
	}
	writerDB.SetMaxOpenConns(10)
	writerDB.SetMaxIdleConns(10)
	// writerDB.SetConnMaxLifetime(time.Second * 1)

	readerDBStmt = &sync.Map{}
	writerDBStmt = &sync.Map{}
}

func SetDataStore(ctx context.Context) context.Context {
	datastore := &DataStore{
		readerDB:    nil,
		writerDB:    nil,
		transaction: nil,
	}
	return context.WithValue(ctx, datastoreKey, datastore)
}

func GetDataStore(ctx context.Context) (*DataStore, error) {
	value := ctx.Value(datastoreKey)
	if value == nil {
		return nil, fmt.Errorf("context datasource not found")
	}
	datasource, ok := value.(*DataStore)
	if !ok {
		return nil, fmt.Errorf("context datasource cast error %v", value)
	}
	return datasource, nil
}

func GetReaderDB(ctx context.Context) (*sql.DB, error) {
	ds, err := GetDataStore(ctx)
	if err != nil {
		return nil, err
	}
	if ds.readerDB == nil {
		ds.readerDB = readerDB
	}
	return ds.readerDB, nil
}

func PrepareReaderDB(ctx context.Context, query string) (*sql.Stmt, error) {
	if stmt, ok := readerDBStmt.Load(query); ok {
		return stmt.(*sql.Stmt), nil
	}

	stmt, err := readerDB.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	readerDBStmt.Store(query, stmt)

	return stmt, nil
}

func PrepareWriterDB(ctx context.Context, query string) (*sql.Stmt, error) {
	if stmt, ok := writerDBStmt.Load(query); ok {
		return stmt.(*sql.Stmt), nil
	}

	stmt, err := writerDB.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	writerDBStmt.Store(query, stmt)

	return stmt, nil
}

func PrepareTx(ctx context.Context, query string) (*sql.Stmt, error) {
	tx, err := GetTransaction(ctx)
	if err != nil {
		return nil, err
	}
	stmt, err := PrepareWriterDB(ctx, query)
	if err != nil {
		return nil, err
	}
	return tx.StmtContext(ctx, stmt), nil
}

func GetTransaction(ctx context.Context) (*sql.Tx, error) {
	ds, err := GetDataStore(ctx)
	if err != nil {
		return nil, err
	}
	if ds.transaction != nil {
		return ds.transaction, nil
	}
	if ds.writerDB == nil {
		ds.writerDB = writerDB
	}
	tx, err := ds.writerDB.Begin()
	if err != nil {
		return nil, err
	}
	ds.transaction = tx
	return ds.transaction, nil
}

func Commit(ctx context.Context) error {
	ds, err := GetDataStore(ctx)
	if err != nil {
		return err
	}
	if ds.transaction != nil {
		err := ds.transaction.Commit()
		if err != nil {
			return err
		}
		ds.transaction = nil
	}
	return nil
}

func Rollback(ctx context.Context) error {
	ds, err := GetDataStore(ctx)
	if err != nil {
		return err
	}
	if ds.transaction != nil {
		err := ds.transaction.Rollback()
		if err != nil {
			return err
		}
		ds.transaction = nil
	}
	return nil
}

func RollbackWithErr(ctx context.Context, err error) error {
	rerr := Rollback(ctx)
	if rerr != nil {
		err = fmt.Errorf(rerr.Error()+": %w", err)
	}
	return err
}
