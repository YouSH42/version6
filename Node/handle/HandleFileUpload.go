package handle

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type MetaDataOne struct {
	Filename string    //`json:"filename"`
	Size     int       //`'json:"size"`
	CreateAT time.Time //`json:"creatat"`
}

func HandleFileUpload(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017") // 코드 가변적으로 변경 가능하게 해보자
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	// 전송된 파일을 가져오는 FormFile
	uploadFile, header, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	defer uploadFile.Close()
	// ***GridFS 생성*** 파일 저장
	db := client.Database("myDB")
	opts := options.GridFSBucket().SetName(header.Filename)
	bucket, err := gridfs.NewBucket(db, opts)
	if err != nil {
		panic(err)
	}
	objectID, err := bucket.UploadFromStream(header.Filename, uploadFile)
	if err != nil {
		panic(err)
	}
	fmt.Printf("New file uploaded with ID %s\n", objectID)
	// 메타데이터 몽고디비에 저장
	// MongoDB 컬렉션 선택
	collection := client.Database("myDB").Collection("metadata")
	err = EnsureCollectionExists(collection)
	if err != nil {
		log.Fatal(err)
	}
	// 메타데이터를 metadata 컬랙션에 저장하는 것
	metadataone := MetaDataOne{
		Filename: header.Filename,
		Size:     int(header.Size),
		CreateAT: time.Now(),
	}
	// 구조체를 BSON 형태로 변환
	bsonData, err := bson.Marshal(metadataone)
	if err != nil {
		log.Fatal(err)
	}
	_, err = collection.InsertOne(context.Background(), bsonData)
	if err != nil {
		log.Fatal(err)
	}

	// 변경: 파일 업로드가 끝난 후 원래 페이지로 리디렉션
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
