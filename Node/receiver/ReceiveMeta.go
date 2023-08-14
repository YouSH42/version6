package receiver

import (
	"fmt"
	"net/http"
	"time"
)

type MetaData struct {
	Filename string
	Size     int
	CreateAT string
	FilePath string
}
type MetaDataOne struct {
	Filename string
	Size     int
	CreateAT time.Time
}
type PageData struct {
	FileSize int
	FileData []MetaDataOne
}
type NodeInfo struct {
	IPAddress string `json:"ip"`
}

var nodes []NodeInfo

// 통신 장치
func ReceiverMeta() {
	http.HandleFunc("/filesave", FileSaveMeta)
	http.HandleFunc("/getinfo", GetInfo)
	http.HandleFunc("/receiveaddr", GetNodeRegist) //노드 정보 받는 api
	http.HandleFunc("/getaddr", GetNodesHandler)   //노드 정보 주는 api
	http.HandleFunc("/clear", ClearNodes)

	fmt.Println("Server listening received ... 2000")
	http.ListenAndServe(":2000", nil)
}
