package webutils

import (
	"log"
	"net/http"
)

func Success(b []byte, w http.ResponseWriter) {
	w.WriteHeader(200)
	writePayload(b, w, false)
}

func NotFound(w http.ResponseWriter) {
	w.WriteHeader(404)
}

func BadRequest(msg string, cause error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusBadRequest)
	writePayload([]byte(msg+"\n"+cause.Error()), w, true)
}

func InternalServerError(msg string, cause error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusInternalServerError)
	writePayload([]byte(msg+"\n"+cause.Error()), w, true)
}

func BadGateway(msg string, cause error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusBadGateway)
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
		log.Printf("Error writing content to http.ResponseWriter: payload=%s, err=%v", payload, err)
	}
}
