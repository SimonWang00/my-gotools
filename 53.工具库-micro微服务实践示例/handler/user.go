package handler

import (
	"context"
	"fmt"
	user_agent "my-gotools/53.工具库-micro微服务实践示例/proto/user"
)

func (s *Service)RpcUserInfo(ctx context.Context,req *user_agent.ReqMsg,res *user_agent.ResMsg)error  {
	fmt.Println(s.userServer.UserInfo(req))
	return nil
}