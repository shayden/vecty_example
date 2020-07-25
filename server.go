//+build ignore
package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./www"))
	log.Fatal(http.ListenAndServe(":5000", fs))
}

func handleRequest(resp http.ResponseWriter, req *http.Request) {

}
