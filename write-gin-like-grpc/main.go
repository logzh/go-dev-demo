package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	svc := &Service{}

	engine := gin.New()

	engine.POST("/users", Wrapper(svc.CreateUser))
	engine.GET("/users", Wrapper(svc.ListUsers))
	engine.GET("/users/:id", Wrapper(svc.GetUser))

	if err := engine.Run(); err != nil {
		log.Fatal(err)
	}
}

type CreateUserRequest struct {
	User *User  `json:"user"`
	Opt  string `form:"opt"`
}

type CreateUserResponse struct {
	User *User  `json:"user"`
	Opt  string `json:"opt"`
}

type ListUsersRequest struct {
	Name string `form:"name"`
}

type ListUsersResponse struct {
	Users []*User `json:"users"`
}

type GetUserRequest struct {
	ID string `form:"id"`
}

type GetUserResponse struct {
	User *User `json:"user"`
}

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Service struct{}

func (s *Service) CreateUser(c *gin.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	// ctx := c.Request.Context()

	return &CreateUserResponse{
		User: req.User,
		Opt:  req.Opt,
	}, nil
}

func (s *Service) ListUsers(c *gin.Context, req *ListUsersRequest) (*ListUsersResponse, error) {
	return &ListUsersResponse{
		Users: []*User{
			{
				ID:    "1",
				Name:  req.Name,
				Email: "min@mail.example.com",
			},
		},
	}, nil
}

func (s *Service) GetUser(c *gin.Context, req *GetUserRequest) (*GetUserResponse, error) {
	if req.ID == "404" {
		return nil, Errorf(404, "user not found")
	}

	if req.ID == "123" {
		return nil, Errorf(123, "user not found")
	}

	return &GetUserResponse{
		User: &User{
			ID:    req.ID,
			Name:  "min",
			Email: "min@mail.example.com",
		},
	}, nil
}

func Errorf(code int, format string, a ...any) error {
	return &apiError{
		error: fmt.Errorf(format, a...),
		code:  code,
	}
}

type APIError interface {
	error

	Code() int
}

type apiError struct {
	error

	code int
}

var _ APIError = &apiError{}

func (e *apiError) Code() int {
	return e.code
}

// Res 统一回复结构
type Res struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// HTTPRsp 统一返回响应处理
func HTTPRsp(c *gin.Context, data interface{}, err error) {
	// ctx := c.Request.Context()
	rsp := Res{}
	rsp.Code = 0
	rsp.Msg = "success"

	tmpErr, ok := err.(*apiError)
	if !ok {
		tmpErr = Errorf(500, "system error").(*apiError)
	}

	// 错误码不为空时，重新赋值
	if err != nil {
		rsp.Code = tmpErr.Code()
		rsp.Msg = tmpErr.Error()
	}

	code := rsp.Code
	switch code {
	case 0:
		rsp.Data = data
		c.PureJSON(http.StatusOK, rsp)
	case 404:
		// https://github.com/gin-gonic/gin/issues/853
		// stop a middleware and get rspponse immediately
		c.AbortWithStatusJSON(http.StatusNotFound, rsp)
	default:
		c.AbortWithStatusJSON(http.StatusAccepted, rsp)
	}
}

// Wrapper ...
func Wrapper[Request, Response any](f func(*gin.Context, *Request) (*Response, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req Request
		var err error
		if c.Request.Method == "GET" {
			err = c.ShouldBindQuery(&req)
		} else {
			err = c.ShouldBindJSON(&req)
		}
		if err != nil {
			HTTPRsp(c, nil, err)
			return
		}

		resp, err := f(c, &req)

		HTTPRsp(c, resp, err)
	}
}
