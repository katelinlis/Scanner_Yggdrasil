package main

import "YggdrasilMap/Default"

/*
type Rec1 struct {
  BoxPubKey string `json:"box_pub_key"`
  //BuildName string `json:"build_name"`
  //BuildVersion string `json:"build_version"`
  Coords string `json:"coords"`
  //Subnet string `json:"subnet"`
}
*/
type Resp1 struct {
	Self map[string]Default.Node `json:"self"`
}

type DataStart struct {
	Request  map[string]interface{} `json:"request"`
	Response Resp1                  `json:"response"`
	Status   string                 `json:"status"`
}

type SettingsConn struct {
	Keepalive bool   `json:"keepalive"`
	Request   string `json:"request"`
}
