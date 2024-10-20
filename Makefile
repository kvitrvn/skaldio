APP_SERVER = server
BUILD_DIR = $(PWD)/bin

.PHONY: test bench run-web

server:
	@go build -o $(BUILD_DIR)/$(APP_SERVER) .

run: server
	@$(BUILD_DIR)/$(APP_SERVER)

test:
	@go test -v ./...

bench:
	@go test -bench=. -benchmem ./...

test-pipeline:
	curl -X POST -H "Content-Type: application/json" --data '{"url":"git@github.com:kvitrvn/skaldio.git", "branch": "main"}' http://127.0.0.1:3000/p/