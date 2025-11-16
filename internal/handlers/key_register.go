package handlers

import (
	"net/http"
	"sayban/internal/models"
)

type KeyRegisterHandler struct {
}

func NewKeyRegisterHadnler() *KeyRegisterHandler {
	return &KeyRegisterHandler{}
}

func (h *KeyRegisterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var keyBundle models.PublicBundle
	err := readJSON(r, &keyBundle)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	err = keyBundle.Validate()
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	writeJSON(w, http.StatusAccepted, map[string]string{
		"response": "successful",
	})
}
