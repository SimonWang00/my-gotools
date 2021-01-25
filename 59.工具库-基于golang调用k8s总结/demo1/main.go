package main

import (
	"fmt"
	core_v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"my-gotools/59.工具库-基于golang调用k8s总结/common"
)

func main() {
	var (
		clientset *kubernetes.Clientset
		podsList *core_v1.PodList
		err error
	)

	// 初始化k8s客户端
	if clientset, err = common.InitClient(); err != nil {
		goto FAIL
	}

	// 获取default命名空间下的所有POD
	if podsList, err = clientset.CoreV1().Pods("default").List(meta_v1.ListOptions{}); err != nil {
		goto FAIL
	}
	fmt.Println(*podsList)

	return

FAIL:
	fmt.Println(err)
	return
}
