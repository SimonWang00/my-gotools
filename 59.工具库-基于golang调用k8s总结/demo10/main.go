package main

import (
	"flag"
	"fmt"
	"k8s.io/client-go/informers"
	core_v1 "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/klog"
	"my-gotools/59.工具库-基于golang调用k8s总结/common"
	"my-gotools/59.工具库-基于golang调用k8s总结/demo10/controller"
	"my-gotools/59.工具库-基于golang调用k8s总结/demo10/pkg/client/clientset/versioned"
	"my-gotools/59.工具库-基于golang调用k8s总结/demo10/pkg/client/informers/externalversions"
	"my-gotools/59.工具库-基于golang调用k8s总结/demo10/pkg/client/informers/externalversions/nginx_controller/v1"
	"time"
)

func main() {
	var (
		restConf *rest.Config
		crdClientset *versioned.Clientset
		clientset *kubernetes.Clientset
		informerFactory informers.SharedInformerFactory
		crdInformerFactory externalversions.SharedInformerFactory
		podInformer core_v1.PodInformer
		nginxInformer v1.NginxInformer
		nginxController *controller.NginxController
		err error
	)

	// 日志参数
	klog.InitFlags(nil)
	flag.Set("logtostderr", "1") // 输出日志到stderr
	flag.Parse()

	// 读取admin.conf, 生成客户端基本配置
	if restConf, err = common.GetRestConf(); err != nil {
		goto FAIL
	}

	// 创建CRD的client
	if crdClientset, err = versioned.NewForConfig(restConf); err != nil {
		goto FAIL
	}

	// 创建K8S内置的client
	if clientset, err = kubernetes.NewForConfig(restConf); err != nil {
		goto FAIL
	}

	// 内建informer工厂
	informerFactory = informers.NewSharedInformerFactory(clientset, time.Second * 120)
	// crd Informer工厂
	crdInformerFactory = externalversions.NewSharedInformerFactory(crdClientset, time.Second * 120)

	// POD informer
	podInformer = informerFactory.Core().V1().Pods()
	// nginx informer
	nginxInformer = crdInformerFactory.Mycompany().V1().Nginxes()

	// 创建调度controller
	nginxController = &controller.NginxController{Clientset: clientset, CrdClientset: crdClientset, PodInformer:podInformer, NginxInformer: nginxInformer}
	nginxController.Start()

	// 等待
	for {
		time.Sleep(1 * time.Second)
	}

	return

FAIL:
	fmt.Println(err)
	return
}
