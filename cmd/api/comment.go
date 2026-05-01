package main

import (
	"fmt"
	"net/http"

	"github.com/Verifieddanny/go-social/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type commentPayload struct {
	Content string `json:"content" validate:"required,max=1000"`
}

func (app *application) createCommentHandler(w http.ResponseWriter, r *http.Request) {
	post := getPostFromCtx(r)

	userId := chi.URLParam(r, "userID")

	if userId == "" {
		app.badRequestResponse(w, r, fmt.Errorf("userID parameter is required"))
		return
	}

	userID, err := uuid.Parse(userId)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	var payload commentPayload

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	comment := &store.Comment{
		PostID:  post.ID,
		UserID:  userID,
		Content: payload.Content,
	}

	ctx := r.Context()

	if err := app.store.Comments.Create(ctx, comment); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, post); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
