package main

import (
	"fmt"
	"net/http"
)

func main() {
	println("server started...")

	http.HandleFunc("/", HelloWorld)
	http.ListenAndServe(":8080", nil)
}
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}
