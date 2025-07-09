package middleware

import (
	"errors"
	"net/http"

	log "github.com/sirupsen/logrus"
	"goapi/api"
	"goapi/internal/tools"
)

var UnauthorizedError = errors.New("invalid username or token")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("username")
		var token string = r.Header.Get("Authorization")
		var err error

		// error & exit if no authorization is provided
		if username == "" || token == "" {
			log.Error(UnauthorizedError)
			api.RequestErrorHandler(w, UnauthorizedError)
			return
		}

		// error & exit if the database has an internal error
		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		// error and exit if the authorization fails
		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username)
		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
			log.Error(UnauthorizedError)
			api.RequestErrorHandler(w, UnauthorizedError)
			return
		}

		/*
			proceed to next middleware or handler
			in this case it will proceed to the GetCoinBalance func
			you can see the next func to be called in internal/handlers/api.go line 18
		*/
		next.ServeHTTP(w, r)
	})
}
