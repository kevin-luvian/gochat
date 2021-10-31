package sample

import (
	"gochat/helper"
	"net/http"

	"github.com/sirupsen/logrus"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	defer logrus.Info("do health check")
	ok := struct {
		Ok bool `json:"ok"`
	}{true}
	helper.SuccessJSON(w, "healthy", ok)
}

func failedCheck(w http.ResponseWriter, r *http.Request) {
	helper.FailedJSON(w, http.StatusBadRequest, "not healthy", nil)
}
