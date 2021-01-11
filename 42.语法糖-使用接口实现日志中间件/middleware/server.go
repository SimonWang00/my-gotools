package middleware

//File  : server.go
//Author: Simon
//Describe: describle your function
//Date  : 2021/1/11

type Service interface {
	Add(a, b int) int
}

type baseServer struct{}

func NewBaseServer() Service {
	return baseServer{}
}

func (s baseServer) Add(a, b int) int {
	return a + b
}

func NewService(s string) Service {
	var svc Service
	{
		svc = NewBaseServer()
		svc = LogMiddleware(s)(svc)
		svc = LogV2Middleware(s)(svc)
	}
	return svc
}
