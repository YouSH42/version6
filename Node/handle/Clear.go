package handle

import (
	"fmt"
	"net/http"
)

func Clear() {
	myaddr := GetAddress()
	url := "http://" + myaddr + ":2000/clear"
	// GET 요청 전송
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error sending request:", err)
	}
	defer resp.Body.Close()

}
