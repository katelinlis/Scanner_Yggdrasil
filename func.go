package main

import (
	"YggdrasilMap/Default"
	"bytes"
	"encoding/json"
	"strings"
	"sync"
	"time"
)

const (
	quotaCount = 10 // Количество горутин
)

func Start() {
	wg := &sync.WaitGroup{}
	quotaChan := make(chan Nodes, quotaCount) // буфферезированный канал(асинхронный), с пустыми структурами(они не занимают места в памяти)

	for i := 1; i <= quotaCount; i++ {
		wg.Add(1)
		//go worker(address, info, t1.Response.Nodes, wg, quotaChan)
		go worker(wg, quotaChan)
	}

	var settingsConn = SettingsConn{true, "getSelf"}
	jsonData, _ := json.Marshal(settingsConn)
	selfInfo := Default.DoRequest(jsonData)

	t1 := DataStart{}
	selfInfo = bytes.Trim(selfInfo, "\x00")
	_ = json.Unmarshal(selfInfo, &t1)

	for ipv6, rec := range t1.Response.Self {
		Gen(ipv6, rec.Coords, Default.DoRequest(Default.GetDHTPingRequest(rec.BoxPubKey, rec.Coords, "")), quotaChan)
		Gen(ipv6, rec.Coords, Default.DoRequest(Default.GetDHTPingRequest(rec.BoxPubKey, rec.Coords, strings.Repeat("0", 128))), quotaChan)
		Gen(ipv6, rec.Coords, Default.DoRequest(Default.GetDHTPingRequest(rec.BoxPubKey, rec.Coords, strings.Repeat("0", 128))), quotaChan)
		wg.Add(1)
	}

	time.Sleep(time.Millisecond)
	wg.Wait()
}

type Nodes struct {
	BoxPubKey  string `json:"box_pub_key"`
	Coords     string `json:"coords"`
	Address    string
	FromIPv6   string
	FromCoords string
}

func Gen(address string, Coords string, data []byte, quotaChan chan Nodes) {
	t1 := DataStart{}
	data = bytes.Trim(data, "\x00")
	_ = json.Unmarshal(data, &t1)

	for ipv6, info := range t1.Response.Self {
		quotaChan <- Nodes{info.BoxPubKey, info.Coords, ipv6, address, Coords}
		PrimaryPeerID := SaveNode(ipv6, Coords)
		SaveLinks(PrimaryPeerID, ipv6, info.Coords)
	}
}
