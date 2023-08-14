package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	bc "test.com/test/Node/broadcast"
	hr "test.com/test/Node/handle"
	sv "test.com/test/Node/receiver"
)

func main() {
	// 노드 생성이 되고 bootstrap서버에 자신의 주소를 넘겨주는 로직
	// address := hr.GetAddress() //자신의 주소를 알아오는 함수
	// hr.RegisterNodeWithBootstrap(address)

	router := httprouter.New()
	router.GET("/", hr.HandleHomeMetaData)
	// router.GET("/test", hr.HandleBroadcast)
	router.POST("/upload/one", hr.HandleFileUpload)
	router.POST("/upload/all", hr.HandleFileUploadAll)

	go bc.Broadcast()
	go sv.ReceiverMeta()

	fmt.Println("server start... localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
