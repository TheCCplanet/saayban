package handlers

import "net/http"

type AccountDeleteHandler struct {
}

func NewAccountDeleteHandler() *AccountDeleteHandler {

	return &AccountDeleteHandler{}
}

func (h *AccountDeleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
