package main

import (
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Errorw("internal server error", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	writeJsonError(w, http.StatusInternalServerError, "The server encountered a problem")

}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warnf("bad request error", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	writeJsonError(w, http.StatusBadRequest, err.Error())
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warnf("not found error", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	writeJsonError(w, http.StatusNotFound, "not found")

}

func (app *application) conflictResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Errorw("conflic error", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	writeJsonError(w, http.StatusConflict, "conflict")

}

func (app *application) unauthorizedResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Errorw("unauthorized error", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	writeJsonError(w, http.StatusUnauthorized, "unauthorized")

}

func (app *application) unauthorizedBasicResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Errorw("unauthorized basic error", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)

	writeJsonError(w, http.StatusUnauthorized, "unauthorized")

}
