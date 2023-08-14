package receiver

import "net/http"

func ClearNodes(_ http.ResponseWriter, _ *http.Request) {
	nodes = nil
}
