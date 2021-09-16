ENV=prod
FUNC=function.$(ENV).json

ifeq ($(ENV),prod)
	FUNC_NAME=inf-prod-rmanzoku
endif

.PHONY: run
run:
	go run main.go

.PHONY: logs
logs:
	lambroll logs --function=$(FUNC) --follow --format="short"

.PHONY: deploy
deploy: handler
	lambroll deploy --function=$(FUNC)

$(FUNC):
	lambroll init --function-name=$(FUNC_NAME)
	mv function.json $@

handler:
	GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o $@