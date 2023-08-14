package receiver

import (
	"net"
	"net/http"
)

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
