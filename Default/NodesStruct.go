package Default

type DataNodes struct {
	Response NodesList `json:"response"`
}
type Node struct {
	BoxPubKey string `json:"box_pub_key"`
	Coords    string `json:"coords"`
}

type NodesList struct {
	Nodes map[string]Node `json:"nodes"`
}
