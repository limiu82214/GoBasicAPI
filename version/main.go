// Package 最基礎的版本，監聽80port並回傳一段文字
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is Basic API\n %s\n", r.URL.Path[1:])
}
