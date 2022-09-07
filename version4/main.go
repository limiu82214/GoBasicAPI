// Package 在這個版本中每個API不局限在main.go裡面，可以歸類於不同的資料夾中
package main

import (
	"fmt"
	"log"
	"net/http"
	"version4/api"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/business/one", api.BusinessOne)
	http.HandleFunc("/business/two", api.BusinessTwo)
	http.HandleFunc("/life/one", api.LifeOne)
	http.HandleFunc("/life/two", api.LifeTwo)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `請嘗試下列網址
localhost/business/one
localhost/business/two
localhost/life/one
localhost/life/two
`)
	fmt.Fprintf(w, "\n")

	fmt.Fprintf(w, "你目前的網址是%v\n", r.URL)
}
