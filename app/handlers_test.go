package app

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"shortlinkapi/database"
	"shortlinkapi/models"
	"testing"
)

type MockCreateUrlInstanceService struct {
	CreateUrlInstanceFunc func(urlInstance models.UrlInstance) error
}

func (m *MockCreateUrlInstanceService) CreateUrlInstance(urlInstance models.UrlInstance) error {
	return m.CreateUrlInstanceFunc(urlInstance)
}

func TestSendUrlHanlder(t *testing.T) {
	t.Run("can create valid url", func(t *testing.T) {
		app := New()
		app.DB = &database.DB{}
		err := app.DB.Open()
		if err != nil {
			t.Fatal(err)
		}
		defer app.DB.Close()

		// counter := 100000000000
		urlInstance := models.UrlInstance{Url: "http://localhost:9000"}

		byteofjson, _ := json.Marshal(urlInstance)
		req, err := http.NewRequest("POST", "/sendurl", bytes.NewReader(byteofjson))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		handler := http.HandlerFunc(app.SendUrlHanlder())
		handler.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Handler returned wrong status code:got %v want %v", status, http.StatusOK)
		}
		var testData map[string]string
		body, err := ioutil.ReadAll(rr.Body)
		err = json.Unmarshal([]byte(body), &testData)
		if err != nil {
			t.Errorf("Decode json has error %v", err)
		}
		var dta map[string]string
		json.Unmarshal([]byte(testData["body"]), &dta)
		if dta["url"] != urlInstance.Url {
			t.Errorf("url is not expected")
		}

	})
}

type testResponse struct {
	message string `json:"message"`
	body    string `json:"body"`
}

func TestIndexHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health-check", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(New().IndexHandler())
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code:got %v want %v", status, http.StatusOK)
	}

	expected := `Welcome to shortlink API`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body : got %v want %v", rr.Body.String(), expected)
	}
}
