// 브로드캐스트 요청이 온 주소로 내 주소를 보내는 함수

package handle

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
)

func SendMyAddrToNode(myaddr string, fromaddr string) {
	// fromaddr 뒤에 포트 번호 때어 내야함
	ip := strings.Split(fromaddr, ":")[0]
	fromAd := "http://" + ip + ":2000/receiveaddr"

	// 생성할 JSON 데이터
	jsonData := []byte(fmt.Sprintf(`{"ip": "%s"}`, myaddr))

	// HTTP POST 요청
	resp, err := http.Post(fromAd, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error sending POST request:", err)
		return
	}
	defer resp.Body.Close()

	// fmt.Println("Response status:", resp.Status)
}
