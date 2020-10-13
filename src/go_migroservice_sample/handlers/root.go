package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Root struct {
	l *log.Logger
}

func NewRoot(l *log.Logger) *Root {
	return &Root{l}
}

func (ro *Root) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ro.l.Println("Root!")

	d, err := ioutil.ReadAll(r.Body)

	if err != nil {
		/* Could do it this way -- left for syntax reference

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("oops"))

		BUT.. */
		http.Error(w, "oops", http.StatusBadRequest)
		return
	}

	ro.l.Printf("Data %s\n", d)
	fmt.Fprintf(w, "Hello! You sent %s in the Root handler request.", d)
}
