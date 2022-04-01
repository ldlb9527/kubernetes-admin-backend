package proto

import "time"

type Labels map[string]string
type Annotations map[string]string

type NameSpace struct {
	Name              string      `json:"name"`
	Status            string      `json:"status"`
	Labels            Labels      `json:"labels"`
	Annotations       Annotations `json:"annotations"`
	CreationTimestamp time.Time   `json:"creation_timestamp"`
}
