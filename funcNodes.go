package main

import (
	"YggdrasilMap/Default"
	"YggdrasilMap/db"
	"YggdrasilMap/db/model"
	"runtime"
	"strings"
	"sync"
)

func SaveNode(address string, Coords string) uint {
	var peer model.Peer
	database := db.Database()
	database.FirstOrCreate(&peer, model.Peer{Coords: Coords, Addr: address})
	_ = database.Close()
	return peer.ID
}

func SaveLinks(PrimaryPeerID uint, ipv6 string, coords string) {
	SecondPeerID := SaveNode(ipv6, coords)
	var peerLink model.PeerLinks
	peerLink.NodeIDPrimary = PrimaryPeerID
	peerLink.NodeIDSecond = SecondPeerID

	database := db.Database()
	database.Save(&peerLink)
	_ = database.Close()
}

func worker(wg *sync.WaitGroup, quotaChan chan Nodes) {
	quotaChan <- Nodes{} // занимаем слот в канале. Если места не будет, то горутина будет ждать и не начнет работу, пока не освободиться место
	defer wg.Done()

	for data := range quotaChan {
		go Gen(data.FromIPv6, data.FromCoords, Default.DoRequest(Default.GetDHTPingRequest(data.BoxPubKey, data.Coords, "")), quotaChan, wg)
		go Gen(data.FromIPv6, data.FromCoords, Default.DoRequest(Default.GetDHTPingRequest(data.BoxPubKey, data.Coords, strings.Repeat("0", 128))), quotaChan, wg)
		go Gen(data.FromIPv6, data.FromCoords, Default.DoRequest(Default.GetDHTPingRequest(data.BoxPubKey, data.Coords, strings.Repeat("f", 128))), quotaChan, wg)
		<-quotaChan
		runtime.Gosched() // передает управление другой горутине
	}
	<-quotaChan // освобождает слот
}
