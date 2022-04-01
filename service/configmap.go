package service

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"kubernetes-admin-backend/client"
	"kubernetes-admin-backend/proto"
	"strconv"
	"time"
)

func ListConfigMap(namespace string) []proto.ConfigMap {
	clientSet, _ := client.GetK8SClientSet()
	configMapList, _ := clientSet.CoreV1().ConfigMaps(namespace).List(context.Background(), metav1.ListOptions{})

	configMaps := make([]proto.ConfigMap, 0, 10)
	for _, item := range configMapList.Items {
		createDays := time.Now().Sub(item.GetCreationTimestamp().Time).Hours() / 24
		configMap := proto.ConfigMap{Name: item.Name, Namespace: item.Namespace, Labels: item.Labels, Annotations: item.Annotations,
			Data: item.Data, Age: strconv.FormatFloat(createDays, 'f', -1, 64) + " days"}
		configMaps = append(configMaps, configMap)
	}
	return configMaps
}

func GetConfigMapByName(namespace, name string) *unstructured.Unstructured {
	dynamicClient, _ := client.GetK8SDynamicClient()
	gvr := schema.GroupVersionResource{Group: "", Version: "v1", Resource: "configmaps"}
	unstructObj, _ := dynamicClient.Resource(gvr).Namespace(namespace).Get(context.Background(), name, metav1.GetOptions{})

	return unstructObj
}
