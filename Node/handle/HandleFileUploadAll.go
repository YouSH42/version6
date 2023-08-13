package handle

import (
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type MetaData struct {
	Filename string
	Size     int
	CreateAT time.Time
	FilePath string
}

func HandleFileUploadAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 웹페이지에서 파일을 가져오는 FormFile
	uploadFile, header, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	defer uploadFile.Close()

	// 메타데이터 설정
	metadata := MetaData{
		Filename: header.Filename,
		Size:     int(header.Size),
		CreateAT: time.Now(),
		FilePath: " ",
	}
	//bootstrap이 되었든 뭐가 되었든 container의 ip와 port를 받아오는 형태로 가야함

	nodes := GetNodeIp()

	// 파일 저장
	for _, node := range nodes {
		uploadFile.Seek(0, 0)
		// fmt.Printf("received IP: %s\n", node.IPAddress)
		FileSaveAll(uploadFile, metadata, node.IPAddress)
	}

	//시간 딜레이가 필요해보임
	time.Sleep(time.Second * 2)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
