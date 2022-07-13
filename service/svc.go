package service

import (
	"context"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubernetes-admin-backend/proto"
)

func ListSvc(clusterName, namespace, label string) []proto.Svc {
	cluster, _ := GetCluster(clusterName)
	clientSet := cluster.ClientSet
	//TODO  将与svc同名的endpoint设置到proto.Svc
	/*endpointsList, _ := clientSet.CoreV1().Endpoints(namespace).List(context.Background(), metav1.ListOptions{LabelSelector: label})
	nameToEndPoints := make(map[string]v1.Endpoints)
	fmt.Println(endpointsList.Items)
	for _, item := range endpointsList.Items {
		item.

	}*/

	serviceList, _ := clientSet.CoreV1().Services(namespace).List(context.Background(), metav1.ListOptions{LabelSelector: label})
	services := serviceList.Items
	svcs := make([]proto.Svc, 0, 10)
	for _, service := range services {
		prots := make([]string, 0, 5)
		for _, servicePort := range service.Spec.Ports {
			if service.Spec.Type == v1.ServiceTypeNodePort {
				s := servicePort.TargetPort.StrVal + ":" + string(servicePort.NodePort) + "/" + string(servicePort.Protocol)
				prots = append(prots, s)
			} else {
				s := servicePort.TargetPort.StrVal + "/" + string(servicePort.Protocol)
				prots = append(prots, s)
			}
		}

		svc := proto.Svc{Name: service.Name, Namespace: service.Namespace,
			Type: string(service.Spec.Type), ClusterIp: service.Spec.ClusterIP,
			Ports: prots, Selector: service.Spec.Selector, Labels: service.Labels, Annotations: service.Annotations,
		}
		svcs = append(svcs, svc)
	}

	return svcs
}
