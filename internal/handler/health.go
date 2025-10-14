package handler

import (
	"net/http"

	"github.com/akhilr007/socials/internal/util"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status": "ok",
	}

	if err := util.WriteJSON(w, http.StatusOK, data); err != nil {
		util.InternalServerError(w, r, err)
	}
}
