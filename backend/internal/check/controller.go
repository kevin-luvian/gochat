package check

import (
	"gochat/helper"
	"net/http"

	"github.com/sirupsen/logrus"
)

type health struct {
	Ok bool `json:"ok"`
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	defer logrus.Info("do health check")
	helper.SuccessJSON(w, "healthy", health{true})
}

func failedCheck(w http.ResponseWriter, r *http.Request) {
	defer logrus.Info("do health check")
	helper.FailedJSON(w, http.StatusBadRequest, "not healthy", health{false})
}
