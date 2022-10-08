package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"shortlinkapi/models"
	"time"
)

func (a *App) IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to shortlink API")
	}
}

func (a *App) SendUrlHanlder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		headerContentType := r.Header.Get("Content-Type")
		if headerContentType != "application/json" {
			errorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
			return
		}
		var urlinstance models.UrlInstance
		var unmarshalErr *json.UnmarshalTypeError

		err := parse(w, r, &urlinstance)
		if err != nil {
			if errors.As(err, &unmarshalErr) {
				errorResponse(w, "Bad Request. Wrong Type provided for field"+unmarshalErr.Field, http.StatusBadRequest)
			} else {
				errorResponse(w, "Bad Request"+err.Error(), http.StatusBadRequest)
			}
			return
		}
		shortLink := &models.ShortUrlLink{
			Url:       urlinstance.Url,
			Count:     0,
			UpdatedAt: time.Now(),
			CreatedAt: time.Now(),
		}
		err = a.DB.CreateShortLink(shortLink)
		if err != nil {
			errorResponse(w, "Insert in DB has error "+err.Error(), http.StatusExpectationFailed)
			return
		}
		errorResponse(w, "Success", http.StatusOK)

	}

}
func parse(w http.ResponseWriter, r *http.Request, data interface{}) error {
	return json.NewDecoder(r.Body).Decode(data)
}
func errorResponse(w http.ResponseWriter, message string, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	resp := make(map[string]string)
	resp["message"] = message
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}
