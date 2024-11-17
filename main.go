package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		http.Error(w, "404 not founnd", http.StatusNotFound)
		return
	}
	if r.Method != "POST" {
		http.Error(w, "method is not suppported", http.StatusNotFound)
		return
	}
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm(): %v", err.Error())
		return
	}
	fmt.Fprintf(w, "POST request successfull")
	fname := r.FormValue("fname")
	lname := r.FormValue("lname")

	fmt.Fprintf(w, "First name: %v", fname)
	fmt.Fprintf(w, "Last name: %v", lname)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method is not suppported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "hellow")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("starting server at port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err.Error())
	}
}
