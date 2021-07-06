package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func uploadHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/upload" {
		http.NotFound(w, r)
		return
	}

	ts, err := template.ParseFiles("./ui/html/index.html")
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal server ERR", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "internal err", 500)
		return
	}
	uploadFile(w, r)

}
