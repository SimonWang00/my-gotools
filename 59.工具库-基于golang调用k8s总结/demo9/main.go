package main

import (
	"fmt"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"my-gotools/59.工具库-基于golang调用k8s总结/common"
	nginx_v1 "my-gotools/59.工具库-基于golang调用k8s总结/demo9/pkg/apis/nginx_controller/v1"
	"my-gotools/59.工具库-基于golang调用k8s总结/demo9/pkg/client/clientset/versioned"
)

func main() {
	var (
		restConf *rest.Config
		crdClientset *versioned.Clientset
		nginx *nginx_v1.Nginx
		err error
	)

	// 读取admin.conf, 生成客户端基本配置
	if restConf, err = common.GetRestConf(); err != nil {
		goto FAIL
	}

	// 创建CRD的client
	if crdClientset, err = versioned.NewForConfig(restConf); err != nil {
		goto FAIL
	}

	// 获取CRD的nginx对象
	if nginx, err = crdClientset.MycompanyV1().Nginxes("default").Get("my-nginx", v1.GetOptions{}); err != nil {
		goto FAIL
	}

	fmt.Println(nginx)

	return

FAIL:
	fmt.Println(err)
	return
}
