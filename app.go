package main

import (
	"PrintHalf/Views"
	"net/http"
)

var (
	route  = views.GetRoute()
	socket = views.GetSocket()
)

func main() {
	//go route.Run(":8080")
	//go socket.Serve()
	//defer socket.Close()
	//http.HandleFunc("/socket.io/", handle)
	////http.Handle("/", http.FileServer(http.Dir("./asset")))
	//log.Println("Serving at localhost:8000...")
	//log.Fatal(http.ListenAndServe(":8000", nil))
	route.Run(":8080")
}

func handle(w http.ResponseWriter, r *http.Request) {
	allowHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Vary", "Origin")
		w.Header().Set("Access-Control-Allow-Methods", "POST, PUT, PATCH, GET, DELETE")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", allowHeaders)
	}
	r.Header.Del("Origin")
	socket.ServeHTTP(w, r)
}
