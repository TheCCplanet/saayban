package handlers

import "net/http"

type AccountRegisterHandler struct {
}

func NewAccountRegisterHandler() *AccountRegisterHandler {

	return &AccountRegisterHandler{}
}

func (h *AccountRegisterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
