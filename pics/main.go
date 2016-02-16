package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/1", getPic1)
	mux.HandleFunc("/2", getPic2)
	mux.HandleFunc("/2.2", getPic22)
	mux.HandleFunc("/3", getPic3)

	fmt.Println("start listening on :8080...")
	fmt.Println(http.ListenAndServe(":8080", mux))
}

func getPic1(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Add("Content-Type", "text/html")
	rw.Write([]byte(floatPic()))
}

func getPic2(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Add("Content-Type", "image/png")
	complexPic1(rw)
}

func getPic22(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Add("Content-Type", "image/png")
	complexPic2(rw)
}
func getPic3(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Add("Content-Type", "image/gif")
	lissajous(rw)
}
