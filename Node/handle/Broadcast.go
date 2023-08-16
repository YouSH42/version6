package handle

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func SendBroadcast() {
	netip := GetAddress()
	parts := strings.Split(netip, ".")
	prefix := ""
	if len(parts) >= 2 {
		prefix = parts[0] + "." + parts[1] + "."
		// fmt.Println(prefix) // 출력: "172.18."
	}

	url := prefix + "255.255:8000"
	// fmt.Println(url)
	conn, err := net.Dial("udp", url) // 브로드캐스트 주소와 포트

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
