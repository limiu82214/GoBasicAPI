// Package 在這個版本中每個API不局限在main.go裡面，可以歸類於不同的資料夾中
package main

import (
	"fmt"
	"log"
	"net/http"
	"version3/api"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // ParseForm之後，r.Form才會有值
	job := r.Form.Get("job")
	method := r.Form.Get("method")
	switch job {
	case "business":
		switch method {
		case "one":
			api.BusinessOne(w, r)
		case "two":
			api.BusinessTwo(w, r)

		}
	case "life":
		switch method {
		case "one":
			api.LifeOne(w, r)
		case "two":
			api.LifeTwo(w, r)

		}
	default:
		fmt.Fprintf(w, `請嘗試下列網址
localhost/?job=business&method=one
localhost/?job=business&method=two
localhost/?job=life&method=one
localhost/?job=life&method=two
`)
		fmt.Fprintf(w, "\n")

		fmt.Fprintf(w, "你目前的網址是%v\n", r.URL)
	}
}
