package repository

import (
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	pb "myshopping/task-srv/proto/task"
	"strings"
	"time"
)

//File  : task.go
//Author: Simon
//Describe: describle your function
//Date  : 2021/3/4

const (
	// 默认数据库名
	DbName = "todolist"
	// 默认表名
	TaskCollection = "task"
	UnFinished     = 0
	Finished       = 1
)

// 定义数据库基本的增删改查操作
type TaskRepository interface {
	InsertOne(ctx context.Context, task *pb.Task) error
	Delete(ctx context.Context, id string) error
	Modify(ctx context.Context, task *pb.Task) error
	Finished(ctx context.Context, task *pb.Task) error
	Count(ctx context.Context, keyword string) (int64, error)
	Search(ctx context.Context, req *pb.SearchRequest) ([]*pb.Task, error)
	// 接口新增方法
	FindById(ctx context.Context, id string) (*pb.Task, error)
}

// 数据库操作实现类
type TaskRepositoryImpl struct {
	// 需要注入一个数据库连接客户端
	Conn *mongo.Client
}

// 定义默认的操作表
func (repo *TaskRepositoryImpl) collection() *mongo.Collection {
	return repo.Conn.Database(DbName).Collection(TaskCollection)
}

// InsertOne 插入一条数据
func (repo *TaskRepositoryImpl) InsertOne(ctx context.Context, task *pb.Task) error {
	_, err := repo.collection().InsertOne(ctx, bson.M{
		"body":       task.Body,
		"startTime":  task.StartTime,
		"endTime":    task.EndTime,
		"isFinished": UnFinished,
		"createTime": time.Now().Unix(),
		// 插入新任务时增加userId
		"userId":     task.UserId,
	})
	return err
}

//Delete 根据id删除数据
func (repo *TaskRepositoryImpl) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = repo.collection().DeleteOne(ctx, bson.M{"_id": oid})
	return err
}

// Modify根据model更新数据
func (repo *TaskRepositoryImpl) Modify(ctx context.Context, task *pb.Task) error {
	id, err := primitive.ObjectIDFromHex(task.Id)
	if err != nil {
		return err
	}
	_, err = repo.collection().UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{
		"body":       task.Body,
		"startTime":  task.StartTime,
		"endTime":    task.EndTime,
		"updateTime": time.Now().Unix(),
	}})
	return err
}

// Finshed更新完成状态
func (repo *TaskRepositoryImpl) Finished(ctx context.Context, task *pb.Task) error {
	// 生成id
	id, err := primitive.ObjectIDFromHex(task.Id)
	if err != nil {
		return err
	}
	now := time.Now().Unix()
	update := bson.M{
		"isFinished": int32(task.IsFinished),
		"updateTime": now,
	}
	if task.IsFinished == Finished {
		update["finishTime"] = now
	}
	log.Print(task)
	log.Println(update)
	_, err = repo.collection().UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": update})
	return err
}

// Count查询符号添加的结果数
func (repo *TaskRepositoryImpl) Count(ctx context.Context, keyword string) (int64, error) {
	filter := bson.M{}
	if keyword != "" && strings.TrimSpace(keyword) != "" {
		filter = bson.M{"body": bson.M{"$regex": keyword}}
	}
	count, err := repo.collection().CountDocuments(ctx, filter)
	return count, err
}

// Search搜索：翻页、排序
func (repo *TaskRepositoryImpl) Search(ctx context.Context, req *pb.SearchRequest) ([]*pb.Task, error) {
	filter := bson.M{}
	if req.Keyword != "" && strings.TrimSpace(req.Keyword) != "" {
		filter = bson.M{"body": bson.M{"$regex": req.Keyword}}
	}

	cursor, err := repo.collection().Find(ctx,
		filter,
		options.Find().SetSkip((req.PageCode-1)*req.PageSize),
		options.Find().SetLimit(req.PageSize),
		options.Find().SetSort(bson.M{req.SortBy: req.Order}))
	if err != nil {
		return nil, errors.WithMessage(err, "search mongo")
	}
	var rows []*pb.Task
	// 解析数据
	if err := cursor.All(ctx, &rows); err != nil {
		return nil, errors.WithMessage(err, "parse data")
	}
	return rows, nil
}


// 通过ID查询task信息
func (repo *TaskRepositoryImpl) FindById(ctx context.Context, id string) (*pb.Task, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.WithMessage(err, "parse ID")
	}
	result := repo.collection().FindOne(ctx, bson.M{"_id": objectId})
	task := &pb.Task{}
	if err := result.Decode(task); err != nil {
		return nil, errors.WithMessage(err, "search mongo")
	}
	return task, nil
}