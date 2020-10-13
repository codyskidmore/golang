package main

import (
	"go_migroservice_sample/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	sm := http.NewServeMux()
	sm.Handle("/", handlers.NewRoot(l))
	sm.Handle("/hello", handlers.NewHello(l))
	sm.Handle("/goodbye", handlers.NewGoodbye(l))

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request){
		log.Println("hello world")
	})

	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request){
		log.Println("goodbye world")
	})

	http.ListenAndServe(":9090", sm)
}

