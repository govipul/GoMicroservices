package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/*
Hello :This is Hello Struct
*/
type Hello struct {
	l *log.Logger
}

//NewHello :new Hello
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello world")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oooooooo", http.StatusBadRequest)
		/*
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("Ooooooooo"))
		*/
		return
	}

	fmt.Fprintf(rw, "Hello %s world \n", d)
}
