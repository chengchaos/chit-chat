package ctrl

import (
	"fmt"
	"net/http"

	"github.com/chengchaos/chit-chat/data"

	"github.com/chengchaos/chit-chat/utils"
)

// NewThread : GET /threads/new
func NewThread(w http.ResponseWriter, r *http.Request) {
	_, err := data.GetSession(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		utils.GenerateHtml(w, nil, "layout", "private.navbar", "new.thread")
	}
}

func CreateThread(w http.ResponseWriter, r *http.Request) {
	sess, err := data.GetSession(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err = r.ParseForm()
		if err != nil {
			utils.LogError(err, "Cannot parse form")
		}
		user, err := data.GetUserBySession(&sess)
		if err != nil {
			utils.LogError(err, "Cannot get user from session")
		}
		topic := r.PostFormValue("topic")
		if _, err := data.CreateThread(&user, topic); err != nil {
			utils.LogError(err, "Cannot create thread")
		}
		http.Redirect(w, r, "/", 302)
	}
}

func PostThread(w http.ResponseWriter, r *http.Request) {
	sess, err := data.GetSession(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err = r.ParseForm()
		if err != nil {
			utils.LogError(err, "Cannot parse form")
		}
		user, err := data.GetUserBySession(&sess)
		if err != nil {
			utils.LogError(err, "Cannot get user from session")
		}
		body := r.PostFormValue("body")
		uuid := r.PostFormValue("uuid")
		thread, err := data.ThreadByUUID(uuid)
		if err != nil {
			utils.ErrorMessage(w, r, "Cannot read thread")
		}
		if _, err := user.CreatePost(thread, body); err != nil {
			utils.LogError(err, "Cannot create post")
		}
		url := fmt.Sprint("/thread/read?id=", uuid)
		http.Redirect(w, r, url, 302)
	}
}

func ReadThread(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	uuid := vals.Get("id")
	thread, err := data.ThreadByUUID(uuid)
	if err != nil {
		utils.ErrorMessage(w, r, "Cannot read thread")
	} else {
		_, err := data.GetSession(w, r)
		if err != nil {
			utils.GenerateHtml(w, &thread, "layout", "public.navbar", "public.thread")
		} else {
			utils.GenerateHtml(w, &thread, "layout", "private.navbar", "private.thread")
		}
	}
}
