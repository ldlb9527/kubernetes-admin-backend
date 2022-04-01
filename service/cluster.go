package service

import (
	"context"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/version"
	"k8s.io/klog"
	"kubernetes-admin-backend/client"
	"kubernetes-admin-backend/proto"
)

func Version() (*version.Info, error) {
	clientSet, err := client.GetK8SClientSet()
	if err != nil {
		klog.Error(err)
		return nil, err
	}
	version, err := clientSet.ServerVersion()
	clientSet.ServerGroups()
	if err != nil {
		klog.Error(err)
		return nil, err
	}
	return version, nil
}

func ListClusters() ([]proto.Node, error) {
	ctx := context.Background()
	clientSet, err := client.GetK8SClientSet()
	if err != nil {
		klog.Error(err)
		return nil, err
	}
	nodeList, err := clientSet.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		klog.Error(err)
		return nil, err
	}

	nodes := make([]proto.Node, 0, 5)
	for _, item := range nodeList.Items {
		node := proto.Node{Name: item.Name, Labels: item.Labels, Annotations: item.Annotations, CreationTimestamp: item.CreationTimestamp.Time,
			Taints: item.Spec.Taints, Status: getReadyStatus(item.Status.Conditions), InternalIp: getInternalIp(item.Status.Addresses),
			KernelVersion: item.Status.NodeInfo.KernelVersion, KubeletVersion: item.Status.NodeInfo.KubeletVersion,
			ContainerRuntimeVersion: item.Status.NodeInfo.ContainerRuntimeVersion, OsImage: item.Status.NodeInfo.OSImage}
		nodes = append(nodes, node)
	}
	return nodes, nil
}

func getInternalIp(addresses []v1.NodeAddress) string {
	for _, address := range addresses {
		if address.Type == v1.NodeInternalIP {
			return address.Address
		}
	}
	return "notfound"
}

func getReadyStatus(conditions []v1.NodeCondition) string {
	for _, condition := range conditions {
		if condition.Type == v1.NodeReady {
			return string(condition.Status)
		}
	}
	return "notfound"
}

func ExtraClusterInfo() proto.ExtraClusterInfo {
	extraClusterInfo := proto.ExtraClusterInfo{0, 0, 0, 0, 0, 0}

	ctx := context.Background()
	clientSet, _ := client.GetK8SClientSet()

	nodeList, _ := clientSet.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	nodes := nodeList.Items
	extraClusterInfo.TotalNodeNum = len(nodes)
	for i := range nodes {
		conditions := nodes[i].Status.Conditions
		for i := range conditions {
			if conditions[i].Type == "Ready" {
				if conditions[i].Status == "True" {
					extraClusterInfo.ReadyNodeNum += 1
				}
			}
		}
		cpu := nodes[i].Status.Allocatable.Cpu().AsApproximateFloat64()
		extraClusterInfo.TotalCpu += cpu
		memory := nodes[i].Status.Allocatable.Memory().AsApproximateFloat64()
		extraClusterInfo.TotalMemory += memory
	}

	podsList, _ := clientSet.CoreV1().Pods("").List(ctx, metav1.ListOptions{})
	pods := podsList.Items
	for i := range pods {
		for j := range pods[i].Spec.Containers {
			cpu := pods[i].Spec.Containers[j].Resources.Requests.Cpu().AsApproximateFloat64()
			extraClusterInfo.UsedCpu += cpu
			memory := pods[i].Spec.Containers[j].Resources.Requests.Memory().AsApproximateFloat64()
			extraClusterInfo.UsedMemory += memory
		}
	}
	return extraClusterInfo
}

func QueryApiGroups() []proto.ApiResource {
	apiResources := make([]proto.ApiResource, 0, 20)
	clientSet, _ := client.GetK8SClientSet()
	_, resources, _ := clientSet.ServerGroupsAndResources()
	for _, resource := range resources {
		for _, api := range resource.APIResources {
			apiResource := proto.ApiResource{Name: api.Name, ShortNames: api.ShortNames, Kind: api.Kind, Namespaced: api.Namespaced, GroupVersion: resource.GroupVersion}
			apiResources = append(apiResources, apiResource)
		}
	}
	return apiResources
}
