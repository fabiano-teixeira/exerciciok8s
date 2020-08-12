package main

import (	
    "fmt"
	"log"
	"net/http"
)

func greeting (sTexto string) string {
	return "<b>" + sTexto + "</b>"
}

func greetingHandler(w http.ResponseWriter, r *http.Request) {

    if r.URL.Path != "/" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }

    if r.Method != "GET" {
        http.Error(w, "Method is not supported.", http.StatusNotFound)
        return
    }
	
	fmt.Fprintf(w, greeting("Code Education Rocks!"))
 }

func main() {

    http.HandleFunc("/", greetingHandler) 

    log.Println ("Servidor Web iniciado na porta 8000.")
    if err := http.ListenAndServe(":8000", nil); err != nil {
        log.Fatal(err)
	}	
}

