package handle

import (
	"context"
	"log"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type PageData struct {
	FileSize int
	FileData []MetaDataOne
}

type NodeFileData struct {
	NodeNumber int
	NodeCount  int
	FileSize   int
	NodeSize   int
	FileData   []MetaDataOne
}

func HandleHomeMetaData(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	//mongoDB연결
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())
	// 컬랙션 선택
	collection := client.Database("myDB").Collection("metadata")
	filter := bson.M{}
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())
	//현재 노드에 있는 리스트
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
	totalnodesize := 0
	var nodeNum PageData
	nodes := GetNodeIp()
	// 받아온 사이즈 정보 저장
	for _, node := range nodes {
		nodeNum, _, _ = ReceiveMetadata(node.IPAddress)
		totalnodesize = totalnodesize + nodeNum.FileSize
	}
	TotalSize = TotalSize / 1000
	totalnodesize = totalnodesize / 1000
	// NodeData에 값 할당
	nodeData := NodeFileData{
		NodeNumber: 0,
		NodeCount:  0,
		FileSize:   TotalSize,
		NodeSize:   totalnodesize,
		FileData:   results,
	}
	tmpl := template.Must(template.ParseFiles("Node/templates/test.html"))
	tmpl.Execute(w, nodeData)
}
