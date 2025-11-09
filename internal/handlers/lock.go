package handlers

import (
	"net/http"
	"sayban/internal/models"
	"sayban/internal/service"
	"strings"
)

type LockHandler struct {
	userService service.UserServiceInterFace
}

func NewLockHandler(userService service.UserServiceInterFace) *LockHandler {
	return &LockHandler{
		userService: userService,
	}
}

func (h *LockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req models.LockRequest

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

	err = h.userService.Lock(req.Name)
	if err != nil {
		// Check if it's an "already locked" error - return 400 Bad Request
		errMsg := err.Error()
		statusCode := http.StatusInternalServerError
		if strings.Contains(errMsg, "already locked") {
			statusCode = http.StatusBadRequest
		} else if strings.Contains(errMsg, "nothing found") {
			statusCode = http.StatusNotFound
		}
		writeError(w, statusCode, err)
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"response": "DB succesfuly locked",
	})
}
