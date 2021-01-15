NAME=lambdagrpcweb

.PHONY: deploy
deploy: $(NAME)
	lambroll deploy

$(NAME): main.go
	GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o $@
