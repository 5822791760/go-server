package main

import (
	"fmt"
	"log"
	"net/http"
)

//====================================================================================================================

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parse Form Error err:%s", err)
		return
	}
	fmt.Fprintf(w, "Form Request Successful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method Not Found", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "HELLO")
}

func serveFiles(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/form.html")
}

//WTFDUDE COLLLdljlsdfkldsjflddlsfjklsdjfkldjsfljsl;dfj;sj;j;fsgj;kojdf;gj;fsdjgo;s;fogji
//sdadsasadsdsa
//sdasdsadas
//====================================================================================================================

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/input", serveFiles)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server at Prot 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
