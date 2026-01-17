.PHONY: all sync lint lint-all lint-gin-demo lint-gin-mcp lint-write-gin-like-grpc

# 默认目标：同步工作区并对所有子项目执行 lint
all: sync lint-all

# 同步 Go 工作区
sync:
	@echo "Syncing Go workspace..."
	@go work sync

# 对所有子项目执行 lint
lint-all:
	@echo "Running lint for all projects..."
	@$(MAKE) lint-gin-demo
	@$(MAKE) lint-gin-mcp
	@$(MAKE) lint-write-gin-like-grpc

# 对单个项目执行 lint
lint: sync lint-all

# gin-demo 项目 lint
lint-gin-demo:
	@echo "Linting gin-demo..."
	@cd gin-demo && gofmt -s -w ./ && goimports -w ./ && golangci-lint run

# gin-mcp 项目 lint
lint-gin-mcp:
	@echo "Linting gin-mcp..."
	@cd gin-mcp && gofmt -s -w ./ && goimports -w ./ && golangci-lint run

# write-gin-like-grpc 项目 lint
lint-write-gin-like-grpc:
	@echo "Linting write-gin-like-grpc..."
	@cd write-gin-like-grpc && gofmt -s -w ./ && goimports -w ./ && golangci-lint run

