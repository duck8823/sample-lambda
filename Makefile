.PHONY: $(shell egrep -o ^[a-zA-Z_-]+: $(MAKEFILE_LIST) | sed 's/://')

help: ## ヘルプを表示する
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z].+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

download: ## モジュールのダウンロード
	@go mod download

generate: ## スキーマからコードを生成する
	@go generate ./...
	@go fmt ./generated/go/user/...

e2e_test: download ## テストを実行する
	@go test main_test.go

lint: download ## Lintを実行する
	@golangci-lint run
