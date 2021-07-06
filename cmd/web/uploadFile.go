package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	//upload max 10MB files
	//parse input from frontEnd
	r.ParseMultipartForm(10 << 20)

	//retrieve files from frntEnd
	// the html input id must match "myFile"
	// <input type="file" name="myFile"/>
	fisierCititHtml, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("err retrieving data")
		fmt.Println(err)
		return
	}
	defer fisierCititHtml.Close()
	fmt.Printf("uploaded file: %v \n", handler.Filename)
	fmt.Printf("file size: %v \n", handler.Size)
	fmt.Printf("MIME header: %v \n", handler.Header)

	//write temp file to serv
	//the imagesDirectory must already be created in root
	tempFile, err := ioutil.TempFile("uploadDirectory", "fileNamePattern-*.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(fisierCititHtml)
	if err != nil {
		fmt.Println(err)
		return
	}
	tempFile.Write(fileBytes)

	//write was success?
	fmt.Fprintf(w, "successfully uploaded file \n")

}
