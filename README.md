# go-dev-demo

一个 Go 开发演示项目集合，包含多个示例项目，用于学习和实践 Go 语言开发、Web 框架使用以及 MCP 协议实现。

## 项目概述

这是一个 Go 工作区（Go Workspace），使用 `go.work` 管理多个相关项目。每个子项目都展示了不同的技术栈和使用场景。

## 项目结构

```
go-dev-demo/
├── gin-demo/              # Gin 框架基础演示
├── gin-mcp/               # Gin + MCP 服务器实现
├── mcp-demo/              # MCP 协议演示
│   └── stdio-demo/        # MCP stdio 传输方式演示
├── write-gin-like-grpc/   # 类似 gRPC 的 Gin 实现
├── go.work                # Go 工作区配置
└── README.md              # 项目文档
```

## 子项目介绍

### gin-demo

Gin 框架的基础演示项目，展示了 Gin Web 框架的核心功能：

- RESTful API 实现
- HTML 模板渲染
- 静态文件服务
- 中间件使用
- JSON 数据处理

**快速开始：**
```bash
cd gin-demo
go run main.go
```

服务将在 `http://localhost:8080` 启动。

### gin-mcp

基于 Gin 框架和 mcp-go 库的 Model Context Protocol (MCP) 服务器实现。

**主要特性：**
- 使用 `StreamableHTTPServer` 实现 MCP 协议
- 支持类型安全的工具定义和处理
- 通过 `http.ServeMux` 挂载，灵活集成到现有路由
- 健康检查端点

**快速开始：**
```bash
cd gin-mcp
go run main.go
```

详细文档请参考 [gin-mcp/README.md](./gin-mcp/README.md)

### mcp-demo/stdio-demo

MCP 协议的 stdio 传输方式演示项目，展示了如何通过标准输入输出实现 MCP 服务器。

**主要特性：**
- 使用 stdio 传输方式
- 计算器工具示例
- 基本的 MCP 协议实现

**快速开始：**
```bash
cd mcp-demo/stdio-demo
go run main.go
```

### write-gin-like-grpc

参考项目，展示了如何用 Gin 框架实现类似 gRPC 的功能。

**参考链接：** https://github.com/min0625/lab-write-gin-like-grpc

## 技术栈

- **Go**: 1.24.11+
- **Gin**: Web 框架
- **mcp-go**: Model Context Protocol Go 实现
- **Go Workspace**: 多模块项目管理

## 环境要求

- Go 1.24.11 或更高版本
- 支持 Go Workspace 功能

## 快速开始

### 1. 克隆项目

```bash
git clone <repository-url>
cd go-dev-demo
```

### 2. 安装依赖

工作区会自动管理所有子项目的依赖，运行：

```bash
go work sync
```

或者进入各个子项目目录分别安装：

```bash
cd gin-demo && go mod download
cd gin-mcp && go mod download
# ... 其他项目
```

### 3. 运行项目

进入对应的子项目目录运行：

```bash
# 运行 Gin 演示
cd gin-demo && go run main.go

# 运行 Gin MCP 服务器
cd gin-mcp && go run main.go

# 运行 MCP stdio 演示
cd mcp-demo/stdio-demo && go run main.go
```

## Go Workspace

本项目使用 Go Workspace 功能管理多个相关项目。工作区配置在 `go.work` 文件中：

```
go 1.24.11

use ./gin-demo
use ./write-gin-like-grpc
use ./mcp-demo/stdio-demo
use ./gin-mcp
```

### Workspace 优势

- **统一管理**：在一个工作区中管理多个相关项目
- **共享依赖**：可以共享和复用依赖包
- **便捷开发**：无需频繁切换目录和模块路径
- **版本一致**：确保所有项目使用相同的 Go 版本

## 开发指南

### 代码格式化

各个子项目都提供了 Makefile，可以使用 `make lint` 进行代码格式化：

```bash
cd gin-demo && make lint
cd gin-mcp && make lint
```

### 添加新项目

1. 在根目录创建新的项目目录
2. 初始化 Go 模块：`go mod init <module-name>`
3. 在 `go.work` 文件中添加 `use ./<project-dir>`
4. 运行 `go work sync` 同步工作区

## 项目用途

这些演示项目可以用于：

- **学习 Go 语言**：了解 Go 的基础语法和特性
- **Web 开发实践**：学习 Gin 框架的使用
- **MCP 协议学习**：理解 Model Context Protocol 的实现
- **项目架构参考**：作为实际项目的参考模板
- **技术栈探索**：尝试不同的技术组合和实现方式

## 相关资源

- [Go 官方文档](https://go.dev/doc/)
- [Gin 框架文档](https://gin-gonic.com/docs/)
- [Model Context Protocol 规范](https://modelcontextprotocol.io/)
- [mcp-go 项目](https://github.com/mark3labs/mcp-go)
- [Go Workspace 文档](https://go.dev/doc/tutorial/workspaces)

## 贡献

欢迎提交 Issue 和 Pull Request 来改进这些演示项目。

## 许可证

各个子项目遵循相应的开源许可证。
