package main

import (
	// "fmt"
	"net"
	"net/http"
	"os"
	"log"
	"io"
)

func main() {
	http.HandleFunc("/", indexFunc)
	http.HandleFunc("/healthz", healthzFunc)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func indexFunc(w http.ResponseWriter, r *http.Request) {
	// 获取request的header，并切入reponse的header中
	for k, v := range r.Header {
		// fmt.Printf("%s:%s\n", k, v)
		w.Header().Set(k, v[0])
	}
	// 获取环境变量VERSION
	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)

	clientIp, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		panic(err)
	}
	statusCode := 200
	if clientIp == "::1" {
		clientIp = "127.0.0.1"
	}
	w.WriteHeader(statusCode)
	log.Printf("[INFO]%d %s\n", statusCode, clientIp)
	io.WriteString(w, "index")
}

func healthzFunc(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	io.WriteString(w, "OK")
}