package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

//File  : main.go
//Author: Simon
//Describe: 连接mongodb使用
//Date  : 2020/12/28

var db *mgo.Session


type User struct {
	Id       bson.ObjectId `bson:"_id"`
	Name     string        `bson:"name"`
	PassWord string        `bson:"pass_word"`
	Age      int           `bson:"age"`
}


func main() {
	// 带登录信息
	dialInfo := &mgo.DialInfo{
		Addrs: 		[]string{"127.0.0.1:27017"}, //远程(或本地)服务器地址及端口号
		Direct: 	false,
		Timeout: 	time.Second * 3,			//超时时间
		Database: 	"admin", 					//数据库
		Source: 	"admin",					//admin
		Username: 	"admin",
		Password: 	"000000",
		PoolLimit: 	4096, 						// Session.SetPoolLimit
	}
	db, err := mgo.DialWithInfo(dialInfo)
	db.SetMode(mgo.Eventual, true)
	if db != nil{
		defer db.Close()
	}
	if err != nil {
		log.Fatalln(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("connect mongodb success!")
	// 配置数据库
	c := db.DB("blog").C("user")
	// 插入测试数据
	func() {
		c.Insert(&User{
			Id:       bson.NewObjectId(),
			Name:     "Simon Wang",
			PassWord: "admin123",
			Age: 2,
		}, &User{
			Id:       bson.NewObjectId(),
			Name:     "Jack Ma",
			PassWord: "666666",
			Age: 5,
		}, &User{
			Id:       bson.NewObjectId(),
			Name:     "Andy Ma",
			PassWord: "888888",
			Age: 7,
		})
	}()

	var users []User
	c.Find(nil).All(&users) //查询全部数据
	log.Println(users)

	c.FindId(users[0].Id).All(&users) //通过ID查询
	log.Println(users)

	c.Find(bson.M{"name": "Jack Ma"}).All(&users) //单条件查询(=)
	log.Println(users)

	c.Find(bson.M{"name": bson.M{"$ne": "Jack Ma"}}).All(&users) //单条件查询(!=)
	log.Println(users)

	c.Find(bson.M{"age": bson.M{"$gt": 5}}).All(&users) //单条件查询(>)
	log.Println(users)

	c.Find(bson.M{"age": bson.M{"$gte": 5}}).All(&users) //单条件查询(>=)
	log.Println(users)

	c.Find(bson.M{"age": bson.M{"$lt": 5}}).All(&users) //单条件查询(<)
	log.Println(users)

	c.Find(bson.M{"age": bson.M{"$lte": 5}}).All(&users) //单条件查询(<=)
	log.Println(users)

	/*c.Find(bson.M{"name": bson.M{"$in": []string{"Jack Ma", "Andy Ma"}}}).All(&users) //单条件查询(in)
	log.Println(users)

	c.Find(bson.M{"$or": []bson.M{bson.M{"name": "Jack Ma"}, bson.M{"age": 7}}}).All(&users) //多条件查询(or)
	log.Println(users)

	c.Update(bson.M{"_id": users[0].Id}, bson.M{"$set": bson.M{"name": "Andy Ma", "age": 61}}) //修改字段的值($set)

	c.FindId(users[0].Id).All(&users)
	log.Println(users)

	c.Find(bson.M{"name": "Andy Ma", "age": 66}).All(&users) //多条件查询(and)
	log.Println(users)

	c.Update(bson.M{"_id": users[0].Id}, bson.M{"$inc": bson.M{"age": -6,}}) //字段增加值($inc)

	c.FindId(users[0].Id).All(&users)
	log.Println(users)*/

	//c.Update(bson.M{"_id": users[0].Id}, bson.M{"$push": bson.M{"interests": "PHP"}}) //从数组中增加一个元素($push)

	c.Update(bson.M{"_id": users[0].Id}, bson.M{"$pull": bson.M{"interests": "PHP"}}) //从数组中删除一个元素($pull)

	c.FindId(users[0].Id).All(&users)
	log.Println(users)

	c.Remove(bson.M{"name": "Jack Ma"})//删除


}


// 连接接口
type SessionStore struct {
	session *mgo.Session
}

//获取数据库的collection
func (d *SessionStore) C(name string) *mgo.Collection {
	return d.session.DB("blog").C(name)
}

// 连接db
func (d *SessionStore) Db() *mgo.Database {
	return d.session.DB("blog")
}

// 可用于异步操作
func NewMgoSession() *SessionStore {
	ds := &SessionStore{
		session: db.Copy(),
	}
	return ds
}

func (d *SessionStore) Close() {
	d.session.Close()
}

func (d *SessionStore) ErrNotFound() error {
	return mgo.ErrNotFound
}

func CloseMgoRedisConnection() {
	db.Close()
}