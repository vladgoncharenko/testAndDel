package main

import (
	"fmt"
	"net/http"
)

//https://docs.google.com/document/d/1V03LUfjSADDooDMhe-_K59EgpTEm3V8uvQRuNMAEnjg/edit

type Handler struct{}

func (*Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, err := Asset("static/gopher.jpg")
	if err != nil {
		fmt.Printf("%v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "image/jpg")
	w.Write(data)
	w.WriteHeader(http.StatusOK)
}

func main() {
	if err := http.ListenAndServe(":8011", new(Handler)); err != nil {
		panic(err)
	}
}

//go:generate go-bindata -o static.go ./static
