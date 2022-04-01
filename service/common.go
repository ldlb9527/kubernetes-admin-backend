package service

import (
	"context"
	"fmt"
	"kubernetes-admin-backend/client"
	//"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/discovery/cached/memory"
	"k8s.io/client-go/restmapper"
	"sigs.k8s.io/yaml"
)

func GetYaml(group, version, resource, namespace, name string) *unstructured.Unstructured {
	dynamicClient, _ := client.GetK8SDynamicClient()
	gvr := schema.GroupVersionResource{Group: group, Version: version, Resource: resource}
	unstructObj, _ := dynamicClient.Resource(gvr).Namespace(namespace).Get(context.Background(), name, metav1.GetOptions{})
	return unstructObj
}

func ApplyYaml(u *unstructured.Unstructured) *unstructured.Unstructured {
	gvk := u.GroupVersionKind()
	gvr, _ := FindGVR(&gvk)

	yamlBytes, _ := yaml.Marshal(u)
	dynamicClient, _ := client.GetK8SDynamicClient()
	//todo 判断不存在则创建,否则更新，细节参考kubectl apply的实现，这里如果存在会报错
	patch, err := dynamicClient.Resource(*gvr).Namespace(u.GetNamespace()).Patch(context.Background(), u.GetName(), types.ApplyPatchType, yamlBytes, metav1.PatchOptions{FieldManager: ""})
	fmt.Println(err)
	return patch
}

// FindGVR 查询gvk对应的gvr
func FindGVR(gvk *schema.GroupVersionKind) (*schema.GroupVersionResource, error) {
	config, _ := client.GetRestConfig()
	dc, err := discovery.NewDiscoveryClientForConfig(config)
	if err != nil {
		return nil, err
	}
	mapper := restmapper.NewDeferredDiscoveryRESTMapper(memory.NewMemCacheClient(dc))
	mapping, err := mapper.RESTMapping(gvk.GroupKind(), gvk.Version)
	if err != nil {
		return nil, err
	}
	//meta.RESTScopeNamespace
	return &mapping.Resource, nil
}
