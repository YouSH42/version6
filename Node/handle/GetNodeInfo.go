package handle

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type NodeInfo struct {
	IPAddress string `json:"ip"`
}

func GetNodeInfo() []NodeInfo {
	// 브로드캐스트로 주소 요청
	SendBroadcast()
	addr := GetAddress()
	// HTTP GET 요청을 보낼 URL
	// bootstrap의 ip주소는 고정이어야 함
	url := "http://" + addr + ":2000/getaddr"
	time.Sleep(time.Second * 2)

	var nodes []NodeInfo

	// GET 요청 전송
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nodes
	}
	defer resp.Body.Close()

	// 응답 읽기
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nodes
	}

	err = json.Unmarshal([]byte(body), &nodes)
	if err != nil {
		log.Fatal(err)
	}

	return nodes
}
