package db

import (
	"github.com/minio/minio-go"
	"my-gotools/31.工具库-独立封装常用数据库、分布式锁、取消信号等/base/config"
	"my-gotools/31.工具库-独立封装常用数据库、分布式锁、取消信号等/base/tool"
)

func initMinio() {
	var secure bool
	if config.GetMinioConfig().GetPath() == "s3.amazonaws.com" {
		secure = true
	}
	if minioClient, err = minio.New(config.GetMinioConfig().GetPath(), config.GetMinioConfig().GetAccessKeyId(), config.GetMinioConfig().GetSecretAccessKey(), secure); err != nil {
		panic(err)
	}
	tool.GetLogger().Debug("50.工具库-minio自动生成缩率图服务 success : " + config.GetMinioConfig().GetPath())
	/*for i:=1;i<=100;i++{
		var bucketName bytes.Buffer
		bucketName.WriteString("storage")
		bucketName.WriteString("-")
		bucketName.WriteString(strconv.Itoa(int(i)))
		fmt.Println(minioClient.MakeBucket(bucketName.String(),""))
	}*/
}
