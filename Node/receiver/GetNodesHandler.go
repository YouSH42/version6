package receiver

import (
	"encoding/json"
	"net/http"
)

func GetNodesHandler(w http.ResponseWriter, _ *http.Request) {
	// 노드 정보를 JSON 형태로 반환
	nodeData, err := json.Marshal(nodes)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(nodeData)
}
