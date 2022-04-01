package proto

type ExtraClusterInfo struct {
	UsedCpu      float64 `json:"used_cpu"`
	TotalCpu     float64 `json:"total_cpu"`
	UsedMemory   float64 `json:"used_memory"`
	TotalMemory  float64 `json:"total_memory"`
	ReadyNodeNum int     `json:"readyNodeNum"`
	TotalNodeNum int     `json:"totalNodeNum"`
}
