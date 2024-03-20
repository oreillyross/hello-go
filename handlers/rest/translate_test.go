package rest_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/oreillyross/hello-go/handlers/rest"
)

func TestTranslateAPI(t *testing.T) {
	tt := []struct {
		EndPoint            string
		StatusCode          int
		ExpectedLanguage    string
		ExpectedTranslation string
	}{
		{
			EndPoint:            "/hello",
			StatusCode:          http.StatusOK,
			ExpectedLanguage:    "english",
			ExpectedTranslation: "hello",
		},
		{
			EndPoint:            "/hello?language=german",
			StatusCode:          http.StatusOK,
			ExpectedLanguage:    "german",
			ExpectedTranslation: "hallo",
		},
		{
			EndPoint:            "/hello?language=dutch",
			StatusCode:          http.StatusNotFound,
			ExpectedLanguage:    "",
			ExpectedTranslation: "",
		},
	}
	// Arrange
	handler := http.HandlerFunc(rest.TranslateHandler)
	for _, test := range tt {
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", test.EndPoint, nil)
		handler.ServeHTTP(rr, req)
		if rr.Code != test.StatusCode {
			t.Errorf(`expected status %d but received %d`, test.StatusCode, rr.Code)

		}
		var resp rest.Resp
		json.Unmarshal(rr.Body.Bytes(), &resp)
		if resp.Language != test.ExpectedLanguage {
			t.Errorf(`Expected %v language but received %v`, test.ExpectedLanguage, resp.Language)
		}
		if resp.Translation != test.ExpectedTranslation {
			t.Errorf(`Expected translation %v but received %v `, test.ExpectedTranslation, resp.Translation)
		}
	}
}
