package repository

//File  : achievement.go
//Author: Simon
//Describe: describle your function
//Date  : 2021/3/4

import (
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

const (
	DbName		   = "todolist"		// 默认数据库名
	TaskCollection = "achievement"	// 默认表名
)

// 数据库的成就记录表结构
type Achievement struct {
	Id string 				`bson:"_id,omitempty"`
	UserId string 			`bson:"userId"`				// 用户ID
	Total int64 			`bson:"total"`				// 完成任务总数
	Finished1Time int64 	`bson:"finished1Time"`		// 完成第一个任务的时间
	Finished100Time int64 	`bson:"finished100Time"`	// 完成第一百个任务的时间
	Finished1000Time int64 	`bson:"finished1000Time"`	// 完成第一千个任务的时间
	UpdateTime int64 		`bson:"updateTime"`			// 更新时间
}

// 成就系统的接口，这里我们定义查询和保存
type AchievementRepo interface {
	FindByUserId(ctx context.Context, userId string) (*Achievement, error)
	Insert(ctx context.Context, achievement *Achievement) error
	Update(ctx context.Context, achievement *Achievement) error
}

// AchievementRepoImpl 接口实现
type AchievementRepoImpl struct {
	Conn *mongo.Client
}

// collection连接库和集合
func (repo *AchievementRepoImpl) collection() *mongo.Collection {
	return repo.Conn.Database(DbName).Collection(TaskCollection)
}

// FindByUserId 根据userid查询
func (repo *AchievementRepoImpl) FindByUserId(ctx context.Context, userId string) (*Achievement, error) {
	result := repo.collection().FindOne(ctx, bson.M{"userId": userId})
	// findOne如果查不到是会报错的,这里要处理一下
	if result.Err() == mongo.ErrNoDocuments {
		return nil, nil
	}
	achievement := &Achievement{}
	if err := result.Decode(achievement); err != nil {
		return nil, errors.WithMessage(err, "search mongo")
	}
	return achievement, nil
}

// Insert 插入1条数据
func (repo *AchievementRepoImpl) Insert(ctx context.Context, achievement *Achievement) error {
	_, err := repo.collection().InsertOne(ctx, achievement)
	return err
}

// 更新
func (repo *AchievementRepoImpl) Update(ctx context.Context, achievement *Achievement) error {
	achievement.UpdateTime = time.Now().Unix()
	oid, err := primitive.ObjectIDFromHex(achievement.Id)
	if err != nil {
		return err
	}
	achievement.Id = ""
	_, err = repo.collection().UpdateOne(ctx,
		bson.M{"_id": oid},
		bson.M{"$set": achievement})
	return err
}