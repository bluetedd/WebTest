

package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"os"
)

var tpl = template.Must(template.ParseFiles("index.html"))

var box1 = "a"
func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)


	var randNum, _ = fmt.Println(rand.Intn(10))

	if  randNum == 5 {
		box1 = "aaaa"
	}


}


