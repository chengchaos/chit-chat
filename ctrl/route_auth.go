package ctrl

import (
	"net/http"

	"github.com/chengchaos/chit-chat/data"

	"github.com/chengchaos/chit-chat/utils"
)

// Login Show the login page
// GET /login
func Login(w http.ResponseWriter, r *http.Request) {
	t := utils.ParseTemplateFiles("login.layout", "public.navbar", "login")
	t.Execute(w, nil)
}

// Signup
// Show the signup page
// GET /sign-up
func Signup(w http.ResponseWriter, r *http.Request) {
	utils.GenerateHtml(w, nil, "login.layout", "public.navbar", "signup")
}

// SignupAccount
// create the user account
// POST /sign-up
func SignupAccount(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		utils.LogError(err, "Connot parse form")
	}
	user := data.User{
		Name:     r.PostFormValue("name"),
		Email:    r.PostFormValue("email"),
		Password: r.PostFormValue("password"),
	}
	if err := user.Create(); err != nil {
		utils.LogError(err, "Cannout create user")
	}
	http.Redirect(w, r, "/login", 302)
}

func Authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	user, err := data.UserByEmail(r.PostFormValue("email"))
	if err != nil {
		utils.LogError(err, "Cannot find user ")
	}
	if user.Password == data.Encrypt(r.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			utils.LogError(err, "Cannot create session")
		}
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("_cookie")
	if err != http.ErrNoCookie {
		utils.LogWarning(err, "Failed to get cookie")
		session := utils.Session{Uuid: cookie.Value}
		data.DeleteSessionByUUID(&session)
	}
	http.Redirect(w, r, "/", 302)
}
