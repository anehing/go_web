package util

import (
	"net/http"
	"github.com/go_web/database"
)

// POST /authenticate
// Authenticate the user given the email and password
func authenticate(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	user, err := database.UserByEmail(request.PostFormValue("email"))
	if err !=nil {
		danger(err, "Cannot find user")
	}

	if user.Password == database.Encrypt(request.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			danger(err, "Cannot create session")
		}
		cookie := http.Cookie{
			Name: "_cookie",
			Value: session.Uuid,
			HttpOnly: true,
		}

		http.SetCookie(writer, &cookie)
		http.Redirect(writer, request, "/", 302)
	} else {
		http.Redirect(writer, request, "/login", 302)
	}
}