package service

import (
	"context"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ListEvent(clusterName, namespace string) *v1.EventList {
	cluster, _ := GetCluster(clusterName)
	clientSet := cluster.ClientSet
	eventList, _ := clientSet.CoreV1().Events(namespace).List(context.Background(), metav1.ListOptions{})
	return eventList
}
