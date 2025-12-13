# Gin MCP Server

一个基于 [Gin](https://github.com/gin-gonic/gin) 框架和 [mcp-go](https://github.com/mark3labs/mcp-go) 的 Model Context Protocol (MCP) 服务器实现。

## 功能特性

- ✅ 基于 Gin 框架，提供高性能的 HTTP 服务
- ✅ 使用 `StreamableHTTPServer` 实现 MCP 协议
- ✅ 支持类型安全的工具定义和处理
- ✅ 通过 `http.ServeMux` 挂载，灵活集成到现有路由
- ✅ 健康检查端点
- ✅ 结构化工具处理，支持输入输出类型定义

## 快速开始

### 前置要求

- Go 1.24.11 或更高版本

### 安装依赖

```bash
go mod download
```

### 运行服务

```bash
go run main.go
```

服务将在 `http://localhost:8080` 启动。

### 使用 Makefile

```bash
# 代码格式化
make lint
```

## 项目结构

```
gin-mcp/
├── main.go          # 主程序入口
├── go.mod           # Go 模块定义
├── go.sum           # 依赖校验和
├── Makefile         # 构建脚本
└── README.md        # 项目文档
```

## API 端点

### MCP 服务端点

- **POST** `/v1/mcp` - MCP 协议端点，处理所有 MCP 请求

### 健康检查

- **GET** `/health` - 健康检查端点，返回服务状态

```bash
curl http://localhost:8080/health
```

响应：
```json
{
  "status": "ok",
  "service": "gin-mcp"
}
```

## 使用示例

### 当前工具

项目包含一个示例工具 `greet`，用于演示如何创建和使用 MCP 工具。

#### Greet 工具

- **名称**: `greet`
- **描述**: 打招呼
- **输入参数**:
  - `name` (string, required): 要问候的名字
- **输出**:
  - `greeting` (string): 问候语

### 如何添加新工具

1. **定义请求和响应结构体**:

```go
type MyToolReq struct {
    Param1 string `json:"param1" jsonschema_description:"参数1" jsonschema:"required"`
    Param2 int    `json:"param2" jsonschema_description:"参数2"`
}

type MyToolRsp struct {
    Result string `json:"result" jsonschema_description:"结果"`
}
```

2. **实现处理函数**:

```go
func myToolHandler(ctx context.Context, req mcp.CallToolRequest, args MyToolReq) (MyToolRsp, error) {
    // 处理逻辑
    return MyToolRsp{Result: "处理结果"}, nil
}
```

3. **注册工具**:

```go
func RegisterMyTool(mcpServer *server.MCPServer) {
    tool := mcp.NewTool("my_tool",
        mcp.WithDescription("工具描述"),
        mcp.WithInputSchema[MyToolReq](),
        mcp.WithOutputSchema[MyToolRsp](),
    )
    mcpServer.AddTool(tool, mcp.NewStructuredToolHandler(myToolHandler))
}
```

4. **在 main 函数中注册**:

```go
RegisterMyTool(mcpServer)
```

## 架构说明

### 集成方式

项目使用以下方式将 MCP 服务集成到 Gin 框架：

```go
// 1. 创建 StreamableHTTPServer
mcpHTTPServer := server.NewStreamableHTTPServer(mcpServer)

// 2. 创建 http.ServeMux 并挂载
mux := http.NewServeMux()
mux.Handle("/v1/mcp", mcpHTTPServer)

// 3. 集成到 Gin 路由
router := gin.Default()
router.POST("/v1/mcp", gin.WrapH(mux))
```

这种方式的优势：
- 保持 MCP 服务器的独立性
- 可以灵活地在 Gin 路由中添加其他端点
- 便于测试和维护

## 依赖项

- `github.com/gin-gonic/gin` - Gin Web 框架
- `github.com/mark3labs/mcp-go` - MCP Go 实现

## 开发

### 代码格式化

```bash
make lint
```

这将运行：
- `gofmt` - Go 代码格式化
- `goimports` - 导入语句格式化
- `golangci-lint` - 代码检查

### 构建

```bash
go build -o gin-mcp main.go
```

## 许可证

本项目遵循相应的开源许可证。

## 相关链接

- [Model Context Protocol 规范](https://modelcontextprotocol.io/)
- [Gin 框架文档](https://gin-gonic.com/docs/)
- [mcp-go 项目](https://github.com/mark3labs/mcp-go)

