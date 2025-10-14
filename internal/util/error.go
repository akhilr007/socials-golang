package util

import (
	"log"
	"net/http"
)

func InternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("internal server error: %s path: %s error: %s", r.Method, r.URL.Path, err)
	WriteJSONError(w, http.StatusInternalServerError, "the server encountered a problem")
}

func BadRequestError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("bad request error: %s path: %s error: %s", r.Method, r.URL.Path, err)
	WriteJSONError(w, http.StatusBadRequest, err.Error())
}

func NotFoundError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("bad request error: %s path: %s error: %s", r.Method, r.URL.Path, err)
	WriteJSONError(w, http.StatusNotFound, "resource not found")
}
