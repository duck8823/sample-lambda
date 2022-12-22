.PHONY: $(shell egrep -o ^[a-zA-Z_-]+: $(MAKEFILE_LIST) | sed 's/://')

help: ## ヘルプを表示する
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z].+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

generate: ## スキーマからコードを生成する
	@go generate ./...
	@go fmt ./generated/go/user/...

test: ## テストを実行する
	@go test ./...
