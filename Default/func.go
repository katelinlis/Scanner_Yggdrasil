package Default

import (
	"encoding/json"
	"net"
)

func GetDHTPingRequest(key string, cords string, target string) []uint8 {
	gettDHTPingRequest1 := DHTPingRequest1{true, "dhtPing", key, cords}
	returnVariable1, _ := json.Marshal(gettDHTPingRequest1)

	if target == "" {
		return returnVariable1
	} else {
		gettDHTPingRequest2 := DHTPingRequest2{true, "dhtPing", key, cords, target}
		returnVariable2, _ := json.Marshal(gettDHTPingRequest2)
		return returnVariable2
	}
}

func DoRequest(jsonData []uint8) []byte {
	conn, err := net.Dial("tcp", "localhost:9001")
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	_, _ = conn.Write(jsonData)
	buffer := make([]byte, 1024*15)
	_, _ = conn.Read(buffer)
	return buffer
}
