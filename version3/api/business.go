package api

import (
	"fmt"
	"net/http"
)

func BusinessOne(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v", "business_one")
}
func BusinessTwo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v", "business_two")
}
