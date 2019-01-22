package route

import (
	"net/http"
	"github.com/go_web/util"
	"github.com/go_web/database"
	"strings"
)

func index(writer http.ResponseWriter, request *http.Request) {
	threads, err := database.Threads()
	if err != nil {
		error_message(writer, request, "Cannot get threads")
	} else {
		_, err := util.Session(writer, request)
		if err != nil {
			util.GenerateHTML(writer, threads, "layout", "public.navbar", "index")
		} else {
			util.GenerateHTML(writer, threads, "layout", "private.navbar", "index")
		}
	}
}

// Convenience function to redirect to the error message page
func error_message(writer http.ResponseWriter, request *http.Request, msg string) {
	url := []string{"/err?msg=", msg}
	http.Redirect(writer, request, strings.Join(url, ""), 302)
}

func err(writer http.ResponseWriter, request *http.Request) {

}