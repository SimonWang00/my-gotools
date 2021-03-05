package subscriber


import (
	"context"
	"github.com/pkg/errors"
	"log"
	pb "myshopping/achievement-srv/proto/task"
	"myshopping/achievement-srv/repository"
	"strings"
	"time"
)

// 定时实现类
type AchievementSub struct {
	Repo repository.AchievementRepo
}

// 只处理任务完成这一个事件
func (sub *AchievementSub) Finished(ctx context.Context, task *pb.Task) error {
	log.Println("Handler Received message: %v\n", task)
	if task.UserId == "" || strings.TrimSpace(task.UserId) == "" {
		return errors.New("userId is blank")
	}
	entity, err := sub.Repo.FindByUserId(ctx, task.UserId)
	if err != nil {
		return err
	}
	now := time.Now().Unix()
	// 查无结果
	if entity == nil {
		entity = &repository.Achievement{
			UserId:        task.UserId,
			Total:         1,
			Finished1Time: now,
		}
		return sub.Repo.Insert(ctx, entity)
	}
	//完成任务总数
	entity.Total++
	// 100 和1000次的时间
	switch entity.Total {
	case 100:
		entity.Finished100Time = now
	case 1000:
		entity.Finished1000Time = now
	}
	return sub.Repo.Update(ctx, entity)

}


// 这个方法保持和Finished方法一致的参数和返回值
func (sub *AchievementSub) Finished2(ctx context.Context, task *pb.Task) error {
	log.Println("Finished2")
	return errors.New("break")
}
// 这个方法去掉了返回值
func (sub *AchievementSub) Finished3(ctx context.Context, task *pb.Task) error{
	log.Println("Finished3")
	return nil
}