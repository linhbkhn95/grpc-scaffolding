
update-gomod:
	go mod tidy

test:
	go test ./...  -count=1 -v -cover -race

test-all-coverage:
	go test ./... -count=1 -race -coverprofile cover.out
	go tool cover -func cover.out

lint: ## Run linter
	golangci-lint run ./...
