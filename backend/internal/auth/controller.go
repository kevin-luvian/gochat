package auth

import (
	"gochat/helper"
	"net/http"
)

func temp(w http.ResponseWriter, r *http.Request) {
	helper.SuccessJSON(w, "healthy", "empty")
}
