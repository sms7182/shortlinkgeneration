package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"shortlinkapi/models"
	"time"

	"github.com/gorilla/mux"
)

const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var UrlShortCode string

func (a *App) IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to shortlink API")
	}
}

func (a *App) RedirectIfURLExist(response http.ResponseWriter, request *http.Request) {
	UrlShortCode = mux.Vars(request)["urlShortCode"]

	if UrlShortCode == "" {
		errorResponse(response, "url not found", http.StatusBadGateway)
		return
	} else {
		actualUrl, err := a.DB.GetActualUrl(UrlShortCode)

		if actualUrl == "" || err != nil {
			errorResponse(response, "actualurl not found and has error", http.StatusBadGateway)
			return
		} else {
			fmt.Print(" url is " + actualUrl)
			http.Redirect(response, request, actualUrl, 302)
		}

	}

}
func (a *App) TestRedirect() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Redirect")
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

		lastcounter, err := a.DB.GetLastCounter()
		if err != nil {
			errorResponse(w, "GetLastCounter has error"+err.Error(), http.StatusExpectationFailed)
			return
		}
		if lastcounter == 0 {
			lastcounter = 100000000000
		}
		shortUrl := longToShort(urlinstance.Url, lastcounter+1)
		shortLink := &models.ShortUrlLink{
			Url:       urlinstance.Url,
			Count:     lastcounter + 1,
			ShortLink: shortUrl,
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

func longToShort(url string, counter int64) string {

	var str string
	n := counter

	for n != 0 {
		str = string(alphabet[n%62]) + str
		n /= 62
	}
	for len(str) != 7 {
		str = string('0') + str
	}
	return str
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
