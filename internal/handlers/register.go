package handlers

import (
	"net/http"
	"sayban/internal/config"
	"sayban/internal/models"
	"sayban/internal/service"
)

type RegisterHandler struct {
	config      *config.Config
	userService service.UserServiceInterFace
}

func NewRegisterHandler(cfg *config.Config, usrService service.UserServiceInterFace) *RegisterHandler {
	return &RegisterHandler{
		config:      cfg,
		userService: usrService,
	}
}

func (h *RegisterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req models.RegisterRequest
	err := readJSON(r, &req)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	err = req.Validate()
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	// create db with and encrypte with their password
	err = h.userService.Register(h.config, req.Name, req.Password)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{
		"response": req.Name + " successfuly was registered",
	})
}
