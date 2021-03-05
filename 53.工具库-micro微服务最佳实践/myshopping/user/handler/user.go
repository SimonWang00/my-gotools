package handler

import (
	"context"
	"github.com/micro/go-micro/errors"

	"github.com/micro/go-micro/util/log"
	"golang.org/x/crypto/bcrypt"

	"myshopping/user/model"
	user "myshopping/user/proto/user"
	"myshopping/user/repository"
)


type User struct{
	UserDb *repository.UserDb
}


// Register 用户注册处理逻辑接口
func (userdb User) Register(ctx context.Context, req *user.RegisterRequest, rsp *user.Response) error  {
	log.Log("Received User.Register request")
	hashpassword, err := bcrypt.GenerateFromPassword([]byte(req.User.Password), bcrypt.DefaultCost)
	if err != nil{
		return err
	}
	muser := &model.User{
		Name:     req.User.Name,
		Phone:    req.User.Phone,
		Password: string(hashpassword),
	}
	if err := userdb.UserDb.Create(muser); err != nil{
		log.Log("register error")
		return err
	}
	rsp.Code = "200"
	rsp.Msg = "恭喜,注册成功!"
	return nil
}

//Login 用户登录处理逻辑
func (userdb User) Login(ctx context.Context, req *user.LoginRequest, rsp *user.Response) error  {
	log.Log("Received User.Login request")
	muser, err := userdb.UserDb.FindByField("phone", req.Phone, "id, password")
	if err != nil{
		return err
	}
	if muser == nil{
		return errors.Unauthorized("go.micro.srv.user.login", "该手机号码不存在")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(muser.Password), []byte(req.Password)); err != nil{
		return errors.Unauthorized("go.micro.srv.user.login", "密码错误")
	}
	rsp.Code = "200"
	rsp.Msg = "登录成功!"
	return nil
}

// UpdatePassword更新用户密码
//流程如下:
//1.先检查是否注册;
//2.再检查旧密码是否正确;
//3.修改密码
func (userdb User)UpdatePassword(ctx context.Context, req *user.UpdatePasswordRequest, rsp *user.Response) error  {
	log.Log("Received User.UpdatePassword request")
	muser, err := userdb.UserDb.Find(req.Uid)
	if muser == nil{
		return errors.Unauthorized("go.micro.srv.user.login", "该用户不存在")
	}
	if err !=  nil{
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(muser.Password), []byte(req.OldPassword)); err != nil{
		return errors.Unauthorized("go.micro.srv.user.login", "旧密码认证失败")
	}
	hashpassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil{
		return err
	}
	muser.Password = string(hashpassword)
	_, err = userdb.UserDb.Update(muser)
	if err != nil{
		return err
	}

	rsp.Code = "200"
	rsp.Msg = muser.Name + ", 您的密码更新成功"
	return nil
}