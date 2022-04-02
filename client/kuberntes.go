package client

import (
	"errors"
	"flag"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/klog"
	"path/filepath"
	"sync"
)

var onceClient sync.Once
var onceDynamicClient sync.Once
var onceConfig sync.Once
var KubeConfig *rest.Config
var KubeClientSet *kubernetes.Clientset

var KubeDynamicClient dynamic.Interface

func GetK8SClientSet() (*kubernetes.Clientset, error) {
	onceClient.Do(func() {
		config, err := GetRestConfig()
		if err != nil {
			klog.Fatal(err)
			return
		}
		KubeClientSet, err = kubernetes.NewForConfig(config)
		if err != nil {
			klog.Fatal(err)
			return
		}
	})
	return KubeClientSet, nil
}

func GetK8SDynamicClient() (dynamic.Interface, error) {
	onceDynamicClient.Do(func() {
		config, err := GetRestConfig()
		if err != nil {
			klog.Fatal(err)
			return
		}
		KubeDynamicClient, err = dynamic.NewForConfig(config)
		if err != nil {
			klog.Fatal(err)
			return
		}
	})
	return KubeDynamicClient, nil
}

func GetRestConfig() (config *rest.Config, err error) {
	onceConfig.Do(func() {
		var kubeConfig *string
		if home := homedir.HomeDir(); home != "" {
			// windows下 home对应 C:\Users\用户名   linux home 对应 /root  这里直接将配置文件放在项目中 不用home
			kubeConfig = flag.String("kubeConfig", filepath.Join("./config", ".kube", "config"), "absolute path to the kubeConfig file")
		} else {
			klog.Fatal("read config error, config is empty")
			err = errors.New("read config error, config is empty")
			return
		}
		flag.Parse()
		KubeConfig, err = clientcmd.BuildConfigFromFlags("", *kubeConfig)
		if err != nil {
			klog.Fatal(err)
			return
		}
		return
	})
	return KubeConfig, nil
}
