package proto

type ApiResource struct {
	Name         string   `json:"name"`
	Kind         string   `json:"kind"`
	ShortNames   []string `json:"short_names"`
	Namespaced   bool     `json:"namespaced"`
	GroupVersion string   `json:"group_version"`
}
