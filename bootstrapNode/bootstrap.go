// bootstrap 노드의 역할:
// 각 노드의 ip정보를 가지고 Get 요청이 들이오면
// 가지고 있는 ip정보 리턴값으로 준다.

package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
)

type NodeInfo struct {
	IPAddress string `json:"ip"`
}

var nodes []NodeInfo

func GetNodesHandler(w http.ResponseWriter, _ *http.Request) {
	// 노드 정보를 JSON 형태로 반환
	nodeData, err := json.Marshal(nodes)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(nodeData)
}

func AddNode(ip string) {
	nodes = append(nodes, NodeInfo{IPAddress: ip})
}

func GetNodeRegist(w http.ResponseWriter, r *http.Request) {
	// TODO: 노드들로 부터 주소값을 받아서 저장
	// 클라이언트 IP 주소 확인
	clientIP, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		http.Error(w, "Invalid client IP address", http.StatusInternalServerError)
		return
	}

	// 받아온 IP 주소를 노드 목록에 추가
	AddNode(clientIP)

	// fmt.Println(clientIP)
	// fmt.Println(what)
}

func main() {
	http.HandleFunc("/getaddr", GetNodesHandler) //노드 정보 주는 api
	http.HandleFunc("/regist", GetNodeRegist)    //노드 정보 받는 api

	fmt.Println("bootstrap listening... 3000")
	http.ListenAndServe(":3000", nil)
}
