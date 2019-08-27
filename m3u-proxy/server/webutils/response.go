package webutils

import (
	"errors"
	"log"
	"net/http"
)

func Success(b []byte, w http.ResponseWriter) {
	w.WriteHeader(200)
	writePayload(b, w, false)
}

func InternalServerError(msg string, cause error, w http.ResponseWriter) {
	w.WriteHeader(500)
	writePayload([]byte(msg+"\n"+cause.Error()), w, true)
}

func BadGateway(msg string, cause error, w http.ResponseWriter) {
	w.WriteHeader(502)
	writePayload([]byte(msg+"\n"+cause.Error()), w, true)
}

// TODO: In the future consider using a struct for this
// It might make the method inputs easier to handle

func writePayload(payload []byte, w http.ResponseWriter, isError bool) {
	if isError {
		log.Printf("An error occured: %s\n", payload)
	}

	_, err := w.Write(payload)
	if err != nil {
		log.Printf(
			"Error writing content to http.ResponseWriter: payload=%s, err=%v",
			payload,
			err)
	}

	errors.New()
}
