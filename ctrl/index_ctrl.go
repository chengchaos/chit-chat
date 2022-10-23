package ctrl

import (
	"net/http"

	"github.com/chengchaos/chit-chat/data"
	"github.com/chengchaos/chit-chat/utils"
)

// Err is a function that shows the error message page
// request : GET /err?msg=
func Err(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	_, err := data.GetSession(w, r)
	if err != nil {
		utils.GenerateHtml(w, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		utils.GenerateHtml(w, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}

func Index(w http.ResponseWriter, r *http.Request) {

	threads, err := data.Threads()
	if err != nil {
		utils.ErrorMessage(w, r, "Cannot get threads")
	} else {
		_, err := data.GetSession(w, r)
		if err != nil {
			utils.GenerateHtml(w, threads, "layout", "public.navbar", "index")
		} else {
			utils.GenerateHtml(w, threads, "layout", "private.navbar", "index")
		}
	}
}
