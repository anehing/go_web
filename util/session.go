package util

import (
	"net/http"
	"github.com/go_web/database"
	"errors"
)

// Checks if the user is logged in and has a session, if not err is not nil
func Session(write http.ResponseWriter, request *http.Request) (session database.Session,err error) {
	cookie, err := request.Cookie("_cookie")
	if err == nil {
		session = database.Session{Uuid: cookie.Value}
		if ok, _ := session.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}
