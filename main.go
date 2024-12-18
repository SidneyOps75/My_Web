package main

import (
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Heelo"))
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is going to be my about page"))
}

func servicesPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my service page"))
}

func formPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(405)
		w.Write([]byte("Method not allowed"))
		return
	}
	w.Write([]byte("Create a new form"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homePage)
	mux.HandleFunc("/about", aboutPage)
	mux.HandleFunc("/services", servicesPage)
	mux.HandleFunc("/form", formPage)

	log.Print("Starting Server at port 8000 ")
	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)
}
