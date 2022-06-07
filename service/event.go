package service

import (
	"context"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubernetes-admin-backend/client"
)

func ListEvent(namespace string) *v1.EventList {
	clientSet, _ := client.GetK8SClientSet()
	eventList, _ := clientSet.CoreV1().Events(namespace).List(context.Background(), metav1.ListOptions{})
	return eventList
}
