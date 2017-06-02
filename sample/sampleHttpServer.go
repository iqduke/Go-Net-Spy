package sample

import (
	"fmt"
	"log"
	"net/http"
)

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello astaxie!")
}

func Run() {

	http.HandleFunc("/", test)

	err := http.ListenAndServe(":8099", nil)

	if err != nil {
		log.Fatal("ListenAndServe:", err)
	} else {
		log.Fatal("ListenAndServe:8099成功，监听中......")
	}
}
