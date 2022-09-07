// Package 在這個版本中以GET的方式接收job和method方法來執行不同的API
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
	r.ParseForm() // ParseForm之後，r.Form才會有值
	job := r.Form.Get("job")
	method := r.Form.Get("method")
	switch job {
	case "business":
		switch method {
		case "one":
			business_one(w, r)
		case "two":
			business_two(w, r)

		}
	case "life":
		switch method {
		case "one":
			life_one(w, r)
		case "two":
			life_two(w, r)

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

func business_one(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v", "business_one")
}
func business_two(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v", "business_two")
}
func life_one(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v", "life_one")
}
func life_two(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v", "life_two")
}
