package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// 创建 MCP 服务器
	mcpServer := server.NewMCPServer(
		"Gin MCP Server",
		"1.0.0",
		server.WithToolCapabilities(false),
	)

	RegisterGreetTool(mcpServer)

	// 创建 StreamableHTTPServer
	mcpHTTPServer := server.NewStreamableHTTPServer(mcpServer)

	// 创建 http.ServeMux 并挂载 MCP 服务器
	mux := http.NewServeMux()
	mux.Handle("/v1/mcp", mcpHTTPServer)

	// 创建 Gin 路由器
	router := gin.Default()

	// 将 ServeMux 集成到 Gin 路由
	// 注意：mux 中已经处理了 /v1/mcp 路径，所以这里直接使用 mux 作为处理器
	router.POST("/v1/mcp", gin.WrapH(mux))

	// 可选：添加健康检查端点
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"service": "gin-mcp",
		})
	})

	// 启动服务器
	fmt.Println("MCP Server starting on :8080")
	if err := router.Run(":8080"); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}

// GreetReq ...
type GreetReq struct {
	Name string `json:"name" jsonschema_description:"name" jsonschema:"required"`
}

// GreetRsp ...
type GreetRsp struct {
	Greeting string `json:"greeting" jsonschema_description:"greeting"`
}

// greetHandler receives typed input and returns typed output
func greetHandler(ctx context.Context, req mcp.CallToolRequest,
	args GreetReq) (GreetRsp, error) {
	// do something
	return GreetRsp{Greeting: fmt.Sprintf("Hello, %s!", args.Name)}, nil
}

// RegisterGreetTool ...
func RegisterGreetTool(mcpServer *server.MCPServer) {
	// Define tool with input and output schemas
	var greetTool = mcp.NewTool("greet",
		mcp.WithDescription("打招呼"),
		mcp.WithInputSchema[GreetReq](),
		mcp.WithOutputSchema[GreetRsp](),
	)

	mcpServer.AddTool(greetTool, mcp.NewStructuredToolHandler(greetHandler))
}
