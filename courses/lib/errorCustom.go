package lib

import (
	"log"
	"net/http"
)

func ReturnInternalError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	if _, err := w.Write([]byte(`Sorry pal, not this time`)); err != nil {
		log.Fatal(err)
	}
}

func ReturnClientError(w http.ResponseWriter, text string) {
	w.WriteHeader(http.StatusBadRequest)
	if _, err := w.Write([]byte(text)); err != nil {
		log.Fatal(err)
	}
}

func LogError(err error) {
	log.Fatal("Error ocurded: ", err)
}

func OnErrFail(e error) {
	if e != nil {
		log.Fatal(e)
	}
}