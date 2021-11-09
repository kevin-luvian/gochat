package helper

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

type JSONResponseBody struct {
	Ok      bool        `json:"ok"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

// wrap responseWriter with header as JSON and status 200.
// Message and object will be marshalled into json in response body.
func SuccessJSON(w http.ResponseWriter, message string, obj interface{}) {
	writeJSONResponse(w, http.StatusOK, true, message, obj)
}

// wrap responseWriter with header as JSON and a dynamic status.
// Message and object will be marshalled into json in response body.
func FailedJSON(w http.ResponseWriter, status int, message string, obj interface{}) {
	writeJSONResponse(w, status, false, message, obj)
}

func writeJSONResponse(w http.ResponseWriter, status int, ok bool, message string, obj interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	res, err := json.Marshal(JSONResponseBody{ok, message, obj})
	if err != nil {
		logrus.Warn("Error when marshalling. ", err.Error())
		logrus.Warn("marshal Obj: ", obj)

		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{ "ok": false, "message": "cannot marshal obj to JSON", result: {} }`)
		return
	}

	w.Write(res)
}
