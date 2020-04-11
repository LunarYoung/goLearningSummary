package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", myHandler);
	http.ListenAndServe("127.0.0.1:8010", nil);
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "Hello World");
	if err != nil {
		return;
	}
	fmt.Println("r.Method =", r.Method);
	fmt.Println("r.URL =", r.URL);
	fmt.Println("r.Header =", r.Header);
	fmt.Println("r.Body =", r.Body);
}
