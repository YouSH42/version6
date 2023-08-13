package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	hr "test.com/test/Node/handle"
	sv "test.com/test/Node/receiver"
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

func main() {
	// 노드 생성이 되고 bootstrap서버에 자신의 주소를 넘겨주는 과정
	address := hr.GetAddress() //자신의 주소를 알아오는 함수
	RegisterNodeWithBootstrap(address)

	router := httprouter.New()
	router.GET("/", hr.HandleHomeMetaData)
	// router.GET("/test", handletest)
	router.POST("/upload/one", hr.HandleFileUpload)
	router.POST("/upload/all", hr.HandleFileUploadAll)

	go sv.ReceiverMeta()

	fmt.Println("server start... localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
