package main

import (
	"fmt"
	"net/http"
)

// ref: https://qiita.com/taizo/items/bf1ec35a65ad5f608d45
// $ ./http_server
// $ curl http://localhost:8080
func main() {
	http.HandleFunc("/", handler) // ハンドラを登録してウェブページを表示させる
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}
