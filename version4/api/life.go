package api

import (
	"fmt"
	"net/http"
)

func LifeOne(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v", "life_one")
}
func LifeTwo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v", "life_two")
}
