package handle

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FileSaveAll(uploadFile multipart.File, metadata MetaData, containerName string) {
	// mongodb Connection
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	// 메타데이터와 파일을 함께 전송하기 위해 multipart/form-data 생성
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	fileWriter, err := writer.CreateFormFile("file", metadata.Filename)
	if err != nil {
		fmt.Println("Error creating form file:", err)
		return
	}
	// 파일 내용을 복사
	_, err = io.Copy(fileWriter, uploadFile)
	if err != nil {
		fmt.Println("Error copying file content:", err)
		return
	}
	// 메타데이터를 JSON 형태로 변환
	jsonData, err := json.Marshal(metadata)
	if err != nil {
		fmt.Println("Error encoding metadata:", err)
		return
	}
	writer.WriteField("metadata", string(jsonData))

	// HTTP POST 요청을 보낼 URL -----1
	url := "http://" + containerName + ":2000/filesave"
	// Content-Type 설정
	writer.Close()
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	// 요청 전송
	clienthttp := &http.Client{}
	resp, err := clienthttp.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Response status:", resp.Status)
}
