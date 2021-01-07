package db

import (
	"github.com/minio/minio-go"
	"my-gotools/all_packaged_library/base/config"
	"my-gotools/all_packaged_library/base/tool"
)

func initMinio() {
	var secure bool
	if config.GetMinioConfig().GetPath() == "s3.amazonaws.com" {
		secure = true
	}
	if minioClient, err = minio.New(config.GetMinioConfig().GetPath(), config.GetMinioConfig().GetAccessKeyId(), config.GetMinioConfig().GetSecretAccessKey(), secure); err != nil {
		panic(err)
	}
	tool.GetLogger().Debug("minio success : " + config.GetMinioConfig().GetPath())
	/*for i:=1;i<=100;i++{
		var bucketName bytes.Buffer
		bucketName.WriteString("storage")
		bucketName.WriteString("-")
		bucketName.WriteString(strconv.Itoa(int(i)))
		fmt.Println(minioClient.MakeBucket(bucketName.String(),""))
	}*/
}
