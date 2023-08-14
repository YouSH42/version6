package handle

import (
	"fmt"
	"net"
	"os"
)

func SendBroadcast() {
	conn, err := net.Dial("udp", "172.18.255.255:8000") // 브로드캐스트 주소와 포트

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	defer conn.Close()

	message := "Hello, world!"
	_, err = conn.Write([]byte(message))

	if err != nil {
		fmt.Println("Error sending message:", err)
		return
	}
}
