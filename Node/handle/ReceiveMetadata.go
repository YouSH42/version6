package handle

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// MetaDataOne이 아니라 PageData를 가져와야함
func ReceiveMetadata(addr string) (PageData, error, bool) {
	var metadatas []MetaDataOne
	nonMetadata := PageData{
		FileSize: 0,
		FileData: metadatas,
	}

	//파일 요청?
	url := "http://" + addr + ":2000/getinfo"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error sending metadata:", err)
		return nonMetadata, fmt.Errorf("error connecting %s", err), false
	}
	defer resp.Body.Close()

	// 응답 데이터 읽기
	var pageData struct {
		FileSize int
		FileData []MetaDataOne
	}
	err = json.NewDecoder(resp.Body).Decode(&pageData)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return nonMetadata, nil, false
	}

	return pageData, nil, true
}
