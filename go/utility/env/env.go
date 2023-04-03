// AWS Secret Manager
package env

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

var sm *secretsmanager.Client

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return
	}
	sm = secretsmanager.NewFromConfig(cfg)
}

func Get(key string) string {
	return GetContext(context.Background(), key)
}

func GetWithDefault(key, defaultValue string) string {
	if Exist(key) {
		return Get(key)
	}
	return defaultValue
}

func GetContext(ctx context.Context, key string) string {
	if ExistEnvVal(key) {
		return os.Getenv(key)
	}

	val, err := GetAWSSecretsManager(ctx, WithPrefix(key))
	if err != nil {
		return ""
	}

	return val
}

func Exist(key string) bool {
	value := Get(key)
	return value != ""
}

func GetAWSSecretsManager(ctx context.Context, key string) (string, error) {
	if sm == nil {
		return "", nil
	}

	params := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(key),
	}
	out, err := sm.GetSecretValue(ctx, params)
	if err != nil {
		return "", err
	}
	return *out.SecretString, nil
}

func WithPrefix(key string) string {
	return os.Getenv("SM_PREFIX") + key
}

func ExistEnvVal(key string) bool {
	return os.Getenv(key) != ""
}
