package main

import (
	"log"
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("internal server error: %s path: %s error: %s", r.Method, r.URL.Path, err)

	writeJsonError(w, http.StatusInternalServerError, "The server encountered a problem")

}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("bad request error: %s path: %s error: %s", r.Method, r.URL.Path, err)

	writeJsonError(w, http.StatusBadRequest, err.Error())

}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("not found error: %s path: %s error: %s", r.Method, r.URL.Path, err)

	writeJsonError(w, http.StatusNotFound, "not found")

}



func (app *application) versionConflictResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("version conflict error: %s path: %s error: %s", r.Method, r.URL.Path, err)

	writeJsonError(w, http.StatusConflict, "version conflict")

}