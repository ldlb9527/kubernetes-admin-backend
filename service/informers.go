package service

import (
	"fmt"
	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog"
	"kubernetes-admin-backend/client"
	"time"
)

func Informers() {
	clientSet, err := client.GetK8SClientSet()
	if err != nil {
		klog.Error(err)
		return
	}
	informerFactory := informers.NewSharedInformerFactory(clientSet, 30*time.Second)
	deployInformer := informerFactory.Apps().V1().Deployments()
	informer := deployInformer.Informer()
	deployLister := deployInformer.Lister()
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    onAdd,
		UpdateFunc: onUpdate,
		DeleteFunc: onDelete,
	})
	stopper := make(chan struct{})
	defer close(stopper)
	//启动Informer List & Watch
	informerFactory.Start(stopper)
	// 等待所有的Informer缓存同步
	informerFactory.WaitForCacheSync(stopper)
	deployments, err := deployLister.Deployments("default").List(labels.Everything())
	//编辑deploy列表
	for index, deployment := range deployments {
		fmt.Printf("%d -> %s\n", index, deployment.Name)
	}
	<-stopper
}

func onAdd(obj interface{}) {
	deployment := obj.(*v1.Deployment)
	klog.Infoln("add a deploy:", deployment)

}

func onUpdate(old, new interface{}) {
	oldDeploy := old.(*v1.Deployment)
	newDeploy := new.(*v1.Deployment)
	klog.Infoln("update deploy:", *oldDeploy.Spec.Replicas, *newDeploy.Spec.Replicas)

}

func onDelete(obj interface{}) {
	deployment := obj.(*v1.Deployment)
	klog.Infoln("delete a deploy:", deployment.Name)
}
