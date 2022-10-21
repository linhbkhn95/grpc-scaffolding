TARGET = bin
TARGET_BIN = rpc-runtime
GO_CMD_MAIN = cmd/main.go

update-gomod:
	go mod tidy

common-update: update-gomod generate

build:
	go build -o $(TARGET)/$(TARGET_BIN) $(GO_CMD_MAIN)

run:
	go run $(GO_CMD_MAIN) server

test:
	go test ./...  -count=1 -v -cover -race

test-all-coverage:
	go test ./... -count=1 -race -coverprofile cover.out
	go tool cover -func cover.out

lint: ## Run linter
	golangci-lint run ./...
