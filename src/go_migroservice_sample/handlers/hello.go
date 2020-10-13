package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.l.Println("hello")

	d, err := ioutil.ReadAll(r.Body)

	if err != nil {
		/* Could do it this way -- left for syntax reference

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("oops"))

		BUT.. */
		http.Error(w, "oops", http.StatusBadRequest)
		return
	}

	h.l.Printf("Data %s\n", d)
	fmt.Fprintf(w, "Hello! You sent %s in the Hello handler request.", d)
}
