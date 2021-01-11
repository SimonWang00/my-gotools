package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	pr "my-gotools/grpc/simple_rpc/proto"
	"net"
)

type Server struct{}

func (s *Server)LoL(ctx context.Context,up *pr.HowieUp)(do *pr.HowieDown,err error) {
	do=&pr.HowieDown{Msg:up.Name+"1111"}
	return do,nil
}


func main() {
	listener,err:=net.Listen("tcp",":8099")
    if err!=nil{
    	log.Panic(err)
	}
	g:=grpc.NewServer()
    pr.RegisterHowieServer(g,&Server{})
	fmt.Println("GRPC 启动成功")
	g.Serve(listener)

}
