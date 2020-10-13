package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g.l.Println("goodbye!")

	d, err := ioutil.ReadAll(r.Body)

	if err != nil {
		/* Could do it this way -- left for syntax reference

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("oops"))

		BUT.. */
		http.Error(w, "oops", http.StatusBadRequest)
		return
	}

	g.l.Printf("Data %s\n", d)
	fmt.Fprintf(w, "Hello! You sent %s in the Goodbye handler request.", d)
}
