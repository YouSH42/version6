package broadcast

import (
	"fmt"
	"net"

	hr "test.com/test/Node/handle"
)

func Broadcast() {
	listenAddr := &net.UDPAddr{IP: net.ParseIP("0.0.0.0"), Port: 8000}

	conn, err := net.ListenUDP("udp", listenAddr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Server listening on broadcast", conn.LocalAddr())

	buffer := make([]byte, 1024)
	for {
		// 클라이언트 메시지 수신 대기
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading message:", err)
			continue
		}

		// 메세지가 날라온 곳으로 자신의 ip를 보내는 함수
		myaddr := hr.GetAddress()
		fromaddr := addr.String()
		// fmt.Println("current address: ", myaddr, "from ", fromaddr)
		hr.SendMyAddrToNode(myaddr, fromaddr)
		// 메세지가 잘 도착했는지 확인
		fmt.Printf("Received %s from %s\n", string(buffer[:n]), addr)
	}
}
