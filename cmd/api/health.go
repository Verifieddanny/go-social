package main

import (
	"net/http"
)

// HealthCheck godoc
//
//	@Summary		health check
//	@Description	health check
//	@Tags			ops
//	@Produce		json
//	@Success		200	{object}	object
//	@Failure		500	{object}	error
//	@Security		ApiKeyAuth
//	@Router			/health [get]
func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":  "ok",
		"env":     app.config.env,
		"version": currentVersion,
	}
	if err := app.jsonResponse(w, http.StatusOK, data); err != nil {
		app.internalServerError(w, r, err)
	}

}
