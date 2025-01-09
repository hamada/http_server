package main

import (
	"net/http"
	"text/template"
)

type Content struct { // テンプレート展開用のデータ構造
	Title string
	Count int
}

func responseHandler(w http.ResponseWriter, r *http.Request) {
	content := Content{"Hello World.", 1}
	tmpl, err := template.ParseFiles("views/layout.html")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, content)
	if err != nil {
		panic(err)
	}
}

// ref: https://qiita.com/taizo/items/bf1ec35a65ad5f608d45
// $ ./http_server
// $ curl http://localhost:8080
func main() {
	http.HandleFunc("/", responseHandler)
	http.ListenAndServe(":8080", nil)
}
