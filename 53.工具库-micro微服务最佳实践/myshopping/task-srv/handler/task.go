package handler

import (
	"context"
	"github.com/micro/go-micro/v2"
	"github.com/pkg/errors"
	"log"
	"time"

	pb "myshopping/task-srv/proto/task"
	"myshopping/task-srv/repository"
)

const(
	// 任务完成消息的topic
	TaskFinishedTopic = "task.finished"
)

type TaskHandler struct{
	repository.TaskRepository
	//TaskRepository repository.TaskRepository
	// 由go-micro封装，用于发送消息的接口，老版本叫micro.Publisher
	TaskFinishedPubEvent micro.Event
}

// Create 创建业务处理逻辑
func (taskhandler *TaskHandler) Create(ctx context.Context, req *pb.Task, rsp *pb.EditResponse) error {
	log.Println("Received TaskSrv.Create request")
	if req.Body == "" || req.StartTime <= 0 || req.EndTime <= 0 || req.UserId == ""{
		return errors.New("bad param")
	}
	if err := taskhandler.TaskRepository.InsertOne(ctx, req); err != nil {
		return err
	}
	rsp.Msg = "create success"
	return nil
}


// Delete 删除任务逻辑
func (taskhandler *TaskHandler) Delete(ctx context.Context, req *pb.Task, rsp *pb.EditResponse) error {
	log.Println("Received TaskSrv.Delete request")
	if req.Id == ""{
		return errors.New("id not valid")
	}
	if err := taskhandler.TaskRepository.Delete(ctx, req.Id); err != nil {
		return err
	}
	rsp.Msg = "已删除id:" + req.Id
	return nil
}


// Modify 更新任务状态
func (taskhandler *TaskHandler) Modify(ctx context.Context, req *pb.Task, rsp *pb.EditResponse) error {
	log.Println("Received TaskSrv.Modify request")
	if req.Id == "" || req.Body == "" || req.StartTime <= 0 || req.EndTime <= 0 {
		return errors.New("bad param")
	}
	if err := taskhandler.TaskRepository.Modify(ctx, req); err != nil {
		return err
	}
	rsp.Msg = "update success"
	return nil
}

// Finished 完成任务状态
func (taskhandler *TaskHandler) Finished(ctx context.Context, req *pb.Task, rsp *pb.EditResponse) error {
	log.Println("Received TaskSrv.Finished request")
	if req.Id == "" || req.IsFinished != repository.UnFinished && req.IsFinished != repository.Finished {
		return errors.New("bad param")
	}
	if err := taskhandler.TaskRepository.Finished(ctx, req); err != nil {
		return err
	}
	rsp.Msg = "success"

	// 发送task完成消息
	// 由于以下都是主业务之外的增强功能，出现异常只记录日志，不影响主业务返回
	if task, err := taskhandler.TaskRepository.FindById(ctx, req.Id); err != nil {
		log.Print("[error]can't send \"task finished\" message. ", err)
	} else {
		if err = taskhandler.TaskFinishedPubEvent.Publish(ctx, task); err != nil {
			log.Print("[error]can't send \"task finished\" message. ", err)
		}
	}
	return nil
}


// Search 根据条件搜索结果
func (taskhandler *TaskHandler) Search(ctx context.Context, req *pb.SearchRequest, rsp *pb.SearchResponse) error {
	log.Println("Received TaskSrv.Search request")
	// 增加3秒延时
	time.Sleep(3 * time.Second)
	count, err := taskhandler.TaskRepository.Count(ctx, req.Keyword)
	if err != nil {
		return errors.WithMessage(err, "count row number")
	}
	if req.PageCode <= 0 {
		req.PageCode = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 20
	}
	if req.SortBy == "" {
		req.SortBy = "createTime"
	}
	if req.Order == 0 {
		req.Order = -1
	}
	if req.PageSize*(req.PageCode-1) > count {
		return errors.New("There's not that much data")
	}
	rows, err := taskhandler.TaskRepository.Search(ctx, req)
	if err != nil {
		return errors.WithMessage(err, "search data")
	}
	*rsp = pb.SearchResponse{
		PageCode: req.PageCode,
		PageSize: req.PageSize,
		SortBy:   req.SortBy,
		Order:    req.Order,
		Rows:     rows,
	}
	return nil
}