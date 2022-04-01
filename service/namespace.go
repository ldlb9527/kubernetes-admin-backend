package service

import (
	"context"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog"
	"kubernetes-admin-backend/client"
	"kubernetes-admin-backend/proto"
)

func GetNamespaces() ([]v1.Namespace, error) {
	ctx := context.Background()
	clientSet, err := client.GetK8SClientSet()
	if err != nil {
		klog.Error(err)
		return nil, err
	}
	namespaceList, err := clientSet.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		klog.Error(err)
		return nil, err
	}
	return namespaceList.Items, nil
}

func CreateNamespace(ns proto.NameSpace) (*v1.Namespace, error) {
	ctx := context.Background()
	clientSet, err := client.GetK8SClientSet()
	if err != nil {
		klog.Error(err)
		return nil, err
	}

	newNamespace, err := clientSet.CoreV1().Namespaces().Create(ctx, &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name:        ns.Name,
			Labels:      ns.Labels,
			Annotations: ns.Annotations,
		},
	}, metav1.CreateOptions{})
	if err != nil {
		klog.Error(err)
		return nil, err
	}

	return newNamespace, nil
}

func DeleteNamespace(nsName string) error {
	ctx := context.Background()
	clientSet, err := client.GetK8SClientSet()
	if err != nil {
		klog.Error(err)
		return err
	}
	deletePolicy := metav1.DeletePropagationForeground
	err = clientSet.CoreV1().Namespaces().Delete(ctx, nsName, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})
	return err
}

func UpdateNamespace(nameSpace proto.NameSpace) (*v1.Namespace, error) {
	ctx := context.Background()
	clientSet, err := client.GetK8SClientSet()
	if err != nil {
		klog.Error(err)
		return nil, err
	}

	namespace := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name:        nameSpace.Name,
			Labels:      nameSpace.Labels,
			Annotations: nameSpace.Annotations,
		},
	}
	update, err := clientSet.CoreV1().Namespaces().Update(ctx, namespace, metav1.UpdateOptions{})
	return update, err
}
