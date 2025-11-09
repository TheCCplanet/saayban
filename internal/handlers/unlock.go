package handlers

import (
	"net/http"
	"sayban/internal/config"
	"sayban/internal/models"
	"sayban/internal/service"
)

type UnlockHandler struct {
	config      *config.Config
	userService service.UserServiceInterFace
}

func NewUnlockHandler(cfg *config.Config, userService service.UserServiceInterFace) *UnlockHandler {
	return &UnlockHandler{
		config:      cfg,
		userService: userService,
	}
}

func (h *UnlockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req models.UnlockRequest

	err := readJSON(r, &req)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
	}

	err = req.Validate()
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	err = h.userService.Unlock(h.config, req.Name, req.Password)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"response": "DB was successfuly unlocked",
	})
}
