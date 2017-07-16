package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/httptest", new(MyHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

type MyHandler func(http.ResponseWriter, *http.Request)

func (self MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	self(w, r)
}
