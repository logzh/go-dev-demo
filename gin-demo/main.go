package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 定义一个简单的用户结构体
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// 模拟数据库中的用户
var users = []User{
	{ID: "1", Username: "user1", Email: "user1@example.com"},
	{ID: "2", Username: "user2", Email: "user2@example.com"},
	{ID: "3", Username: "user3", Email: "user3@example.com"},
}

func main() {
	// 创建默认的Gin路由器
	r := gin.Default()

	// 加载HTML模板
	r.LoadHTMLGlob("templates/*")

	// 提供静态文件服务
	r.Static("/static", "./static")

	// 定义一个自定义中间件
	r.Use(customMiddleware())

	// 首页路由
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Gin框架演示",
		})
	})

	// API路由组
	api := r.Group("/api")
	{
		// 获取所有用户
		api.GET("/users", getUsers)
		
		// 获取单个用户
		api.GET("/users/:id", getUserByID)
		
		// 创建用户
		api.POST("/users", createUser)
	}

	// 启动服务器
	r.Run(":8080") // 监听并在0.0.0.0:8080上启动服务
}

// 自定义中间件
func customMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 请求前
		c.Set("example", "middleware value")

		c.Next()

		// 请求后
		// 可以在这里添加一些请求后的逻辑
	}
}

// 获取所有用户的处理函数
func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

// 根据ID获取用户的处理函数
func getUserByID(c *gin.Context) {
	id := c.Param("id")
	
	for _, user := range users {
		if user.ID == id {
			c.JSON(http.StatusOK, gin.H{
				"user": user,
			})
			return
		}
	}
	
	c.JSON(http.StatusNotFound, gin.H{
		"message": "用户不存在",
	})
}

// 创建用户的处理函数
func createUser(c *gin.Context) {
	var newUser User
	
	// 绑定JSON请求体到User结构体
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	// 简单地将新用户添加到用户列表中
	// 在实际应用中，这里会与数据库交互
	users = append(users, newUser)
	
	c.JSON(http.StatusCreated, gin.H{
		"message": "用户创建成功",
		"user":    newUser,
	})
}