package main

import (
	"net/http"
	"os"

	"github.com/berryhe/cloud-native-curriculum/week_01/transport"
)

func main() {
	transport.Version = os.Getenv("VERSION")

	http.HandleFunc("/", transport.HandleRootPath)
	http.HandleFunc("/healthz", transport.HandleHealthz)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
