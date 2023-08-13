// 노드 정보 불러오는 코드
package receiver

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

func GetInfo(w http.ResponseWriter, _ *http.Request) {
	// mongoDB에서 메타데이터 가져오기
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())
	// 컬랙션 선택
	collection := client.Database("myDB").Collection("metadata")
	// 컬랙션 내에 모든 파일 검사
	filter := bson.M{}
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())
	// 메타데이터 저장
	var results []MetaDataOne
	TotalSize := 0
	for cur.Next(context.Background()) {
		var data MetaDataOne
		if err := cur.Decode(&data); err != nil {
			log.Fatal(err)
		}
		TotalSize = TotalSize + data.Size
		results = append(results, data)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	// PageData에 값 할당
	pageData := PageData{
		FileSize: TotalSize,
		FileData: results,
	}
	// JSON 형태로 인코딩
	jsonData, err := json.Marshal(pageData)
	if err != nil {
		http.Error(w, "Error encoding data", http.StatusInternalServerError)
		return
	}

	// 응답 전송
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
