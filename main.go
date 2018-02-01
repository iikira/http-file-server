package main

import (
	"flag"
	"fmt"
	"github.com/iikira/BaiduPCS-Go/pcsutil"
	"net"
	"net/http"
	"net/url"
)

var (
	port uint
	path string
)

func init() {
	flag.UintVar(&port, "port", 80, "port")
	flag.StringVar(&path, "path", "/", "local relative or absolute path")

	flag.Parse()
}

func main() {
	http.Handle("/", http.FileServer(http.Dir(path)))

	fmt.Println("Server is starting...")
	// Print available URLs.
	for _, address := range pcsutil.ListAddresses() {
		fmt.Printf("URL: %s\n", (&url.URL{
			Scheme: "http",
			Host:   net.JoinHostPort(address, fmt.Sprint(port)),
			Path:   "/",
		}).String())
	}
	fmt.Println("ListenAndServe:", http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
