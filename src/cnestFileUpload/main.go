package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func uploadHandler(rw http.ResponseWriter, r *http.Request) {

	uploadFile, header, err := r.FormFile("upload_file")	
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(rw, err)
		return
	}
	defer uploadFile.Close()

	dirname := "uploads"
	os.MkdirAll(dirname, 0777)
	
	filepath := fmt.Sprintf("%s/%s", dirname, header.Filename)
	file, err := os.Create(filepath) // 빈파일을 만듬	
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(rw, err)
		return
	}
	defer file.Close()

	io.Copy(file, uploadFile)
	rw.WriteHeader(http.StatusOK)
	fmt.Fprint(rw, filepath)
}


func main() {
	http.HandleFunc("/uploads", uploadHandler)
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":3000", nil)
}