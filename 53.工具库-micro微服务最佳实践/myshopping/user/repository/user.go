package repository

//File  : user.go
//Author: Simon
//Describe: describle your function
//Date  : 2021/3/3

import (
	"github.com/jinzhu/gorm"
	"myshopping/user/model"
)

type Repository interface{
	Find(id int32) (*model.User, error)
	Create(*model.User) error
	Update(*model.User, int64) (*model.User, error)
	FindByField(string, string, string) (*model.User, error)
}

type UserDb struct {
	Db *gorm.DB
}

//Find 根据id查询用户
func (userdb *UserDb) Find(id uint32) (*model.User, error) {
	user := &model.User{}
	user.ID = uint(id)
	if err := userdb.Db.First(user).Error; err != nil{
		return nil, err
	}
	return user, nil
}

//Create 注册用户
func (userdb UserDb) Create(user *model.User) error {
	if err := userdb.Db.Create(user).Error; err != nil{
		return err
	}
	return nil
}

//Update 更新资料
func (userdb UserDb) Update(user *model.User) (*model.User, error)  {
	if err := userdb.Db.Model(user).Updates(&user).Error; err != nil{
		return nil, err
	}
	return user, nil
}

//根据字段查找
func (userdb UserDb) FindByField(key string, value string, fields string) (*model.User, error)  {
	if len(fields) == 0{
		fields = "*"
	}
	user := &model.User{}
	if err := userdb.Db.Select(fields).Where(key + " = ?",value).First(&user).Error; err != nil{
		return nil, err
	}
	return user, nil
}