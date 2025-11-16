package handlers

import "net/http"

type AccountListHandler struct {
}

func NewAccountListHandler() *AccountListHandler {

	return &AccountListHandler{}
}

func (h *AccountListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
