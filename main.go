package main

import (
	"fmt"
	"net/http"
)

func main() {
	println("server started...")

	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8080", nil)
}
func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}
