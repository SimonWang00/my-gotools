package weblib

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	grpcSvc "myshopping/client/proto/user"
)

//File  : router.go
//Author: Simon
//Describe: describle your function
//Date  : 2021/3/4

func NewGinRouter(prodSvc grpcSvc.UserService) *gin.Engine {
	ginRouter := gin.Default() //gin web框架
	ginRouter.Use(InitMiddleware(prodSvc))
	v1Group := ginRouter.Group("/v1")
	{
		v1Group.Handle("POST", "/prods", Mydemo) //组成 POST /v1/prods 路由
	}

	return ginRouter
}

//middleware
//为了将grpccli封装到context中
func InitMiddleware(prodSvc grpcSvc.UserService) gin.HandlerFunc  {
	return func(context *gin.Context) {
		context.Keys = make(map[string]interface{})
		context.Keys["userservice"] = prodSvc

		context.Next()
	}
}
//统一错误处理
func ErrorMiddleware() gin.HandlerFunc  {
	return func(context *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				context.JSON(500, gin.H{"status":fmt.Sprintf("%s", r)})
				context.Abort()
			}
			context.Next()
		}()
	}
}

//handler
func defaultProds() (resp *grpcSvc.Response, err error) {
	resp = &grpcSvc.Response{Code:"200", Msg:"register ok"} //调用grpc server
	return
}


func PanicError(err error)  {
	if err != nil {
		panic(err)
	}
}

func GetProdDetail(ginCtx *gin.Context)  {
	var prodReq grpcSvc.RegisterRequest //此pb结构体添加注解，表示要支持uri可变参数
	ginCtx.JSON(200, gin.H{"data": prodReq.User.Name})
	return
}

func Mydemo(ginCtx *gin.Context)  {
	var req grpcSvc.RegisterRequest
	var resp *grpcSvc.Response

	prodSvc := ginCtx.Keys["userservice"].(grpcSvc.UserService)
	err := ginCtx.Bind(&req)
	if err != nil {
		ginCtx.JSON(500, gin.H{"status":err.Error()})
	} else {
		resp, err =  prodSvc.Register(context.Background(), &req)
		if err != nil{
			print("Register error," )
			panic( err)
		}
		ginCtx.JSON(200, gin.H{"code": resp.Code, "msg":resp.Msg})
		return
	}
}