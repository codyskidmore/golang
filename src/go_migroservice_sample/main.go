package main

import (
	"context"
	"go_migroservice_sample/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
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

	s := &http.Server{
		Addr: ":9090",
		Handler: sm,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <- sigChan
	l.Println("received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30 * time.Second)
	s.Shutdown(tc)
}

