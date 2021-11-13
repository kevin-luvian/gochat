package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

func RequestBody(r *http.Request) (map[string]interface{}, error) {
	m := map[string]interface{}{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.Error("Error reading body: ", err)
		return nil, err
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, &m); err != nil {
		logrus.Error("Failed to unmarshal ", err)
		return nil, err
	}
	return m, nil
}
