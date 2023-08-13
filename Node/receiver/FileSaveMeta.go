// 파일 전체 업로드
package receiver

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
	hr "test.com/test/Node/handle"
)

func FileSaveMeta(w http.ResponseWriter, r *http.Request) {
	// 파일 파트 가져오기
	filePart, _, err := r.FormFile("file")
	if err != nil {
		// 에러 처리
		return
	}
	defer filePart.Close()
	// fmt.Println("real file name: ", fileHeader.Filename)
	// fmt.Println("real file size: ", fileHeader.Size)

	// 메타데이터 파트 가져오기
	metadataPart := r.FormValue("metadata")

	//mongoDB 연결
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017") // 코드 가변적으로 변경 가능하게 해보자
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	// 메타데이터 수신 string 타입을 json(구조체)로 변환
	var metadata MetaData
	err = json.Unmarshal([]byte(metadataPart), &metadata)
	if err != nil {
		fmt.Println("Error converting JSON:", err)
		return
	}
	// ***GridFS 생성***
	// database 지정(없을 시 새로 만듬)
	db := client.Database("myDB")
	opts := options.GridFSBucket().SetName(metadata.Filename)
	bucket, err := gridfs.NewBucket(db, opts)
	if err != nil {
		panic(err)
	}

	objectID, err := bucket.UploadFromStream(metadata.Filename, filePart)
	if err != nil {
		panic(err)
	}
	fmt.Printf("New file uploaded with ID %s\n", objectID)
	// 메타데이터 몽고디비에 저장
	// MongoDB 컬렉션 선택
	collection := client.Database("myDB").Collection("metadata")
	err = hr.EnsureCollectionExists(collection)
	if err != nil {
		log.Fatal(err)
	}
	// 메타데이터를 metadata 컬랙션에 저장하는 것
	metadataone := MetaDataOne{
		Filename: metadata.Filename,
		Size:     int(metadata.Size),
		CreateAT: time.Now(),
	}
	// fmt.Println("metadata filename: ", metadataone.Filename)
	// fmt.Println("metadata size: ", metadataone.Size)
	// 구조체를 BSON 형태로 변환
	bsonData, err := bson.Marshal(metadataone)
	if err != nil {
		log.Fatal(err)
	}
	_, err = collection.InsertOne(context.Background(), bsonData)
	if err != nil {
		log.Fatal(err)
	}
	// 응답 보내기
	w.Write([]byte("Upload successful"))
}
