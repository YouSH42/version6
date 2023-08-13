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

// 통신 장치
func ReceiverMeta() {
	http.HandleFunc("/metadata", FileSaveMeta)
	http.HandleFunc("/getinfo", GetInfo)

	fmt.Println("Server listening... 2000")
	http.ListenAndServe(":2000", nil)
}
