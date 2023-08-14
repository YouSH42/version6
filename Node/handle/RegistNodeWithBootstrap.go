package handle

import (
	"bytes"
	"fmt"
	"net/http"
)

func RegisterNodeWithBootstrap(nodeIP string) error {
	// bootstrap 서버의 주소
	bootstrapURL := "http://bootstrap:3000/regist"

	// POST 요청을 위한 JSON 데이터 생성
	requestData := []byte(fmt.Sprintf(`{"ip": "%s"}`, nodeIP))

	// POST 요청 보내기
	resp, err := http.Post(bootstrapURL, "application/json", bytes.NewBuffer(requestData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 응답 코드 확인
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to register node with bootstrap: %s", resp.Status)
	}

	// fmt.Println("Node registered with bootstrap:", nodeIP)
	return nil
}
