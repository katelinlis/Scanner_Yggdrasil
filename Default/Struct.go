package Default

type DHTPingRequest1 struct {
	Keepalive bool   `json:"keepalive"`
	Request   string `json:"request"`
	BoxPubKey string `json:"box_pub_key"`
	Coords    string `json:"coords"`
}
type DHTPingRequest2 struct {
	Keepalive bool   `json:"keepalive"`
	Request   string `json:"request"`
	BoxPubKey string `json:"box_pub_key"`
	Coords    string `json:"coords"`
	Target    string `json:"target"`
}
