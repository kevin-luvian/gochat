package auth

import (
	"gochat/env"
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	env.LoadDotEnvForTest()
}

func TestLoginWithGmailRedirect(t *testing.T) {
	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(loginGoogle)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusTemporaryRedirect {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusTemporaryRedirect)
	}
}
