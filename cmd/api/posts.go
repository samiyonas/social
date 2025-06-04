package main

import (
	"net/http"
)

type createPayload struct {
	Title string `json:"title"`
	Content string `json:"content"`
	tags []string `json:"tags"`
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	userID := 1

	var payload createPayload

	if err := readJSON(w, r, &payload); err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	post := &store.Post{
		title: payload.title,
		content: payload.content,
		tags: payload.Tags,
		UserID: payload.UserID,
	}

	ctx := r.Context()

	if err := app.store.post.Create(ctx, post); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := writeJSON(w, http.StatusOk, post); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
}
