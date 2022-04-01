package proto

type ConfigMap struct {
	Age         string            `json:"age"`
	Name        string            `json:"name"`
	Data        map[string]string `json:"data"`
	Labels      Labels            `json:"labels"`
	Namespace   string            `json:"namespace"`
	Annotations Annotations       `json:"annotations"`
}
