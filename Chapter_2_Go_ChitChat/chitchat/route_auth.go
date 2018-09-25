package main

import (
	"./data"
	"net/http"
)

// GET /login
// Show the login page
func login(writer http.ResponseWriter, request *http.Request) {
	t := parseTemplateFiles("login.layout", "public.navbar", "login")
	t.Execute(writer, nil)
}

// GET /signup
// Show the signup page
func signup(writer http.ResponseWriter, request *http.Request) {
	generateHTML(writer, nil, "login.layout", "public.navbar", "signup")
}

// POST /signup
// Create the user account
func signupAccount(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		danger(err, "Cannot parse form")
	}
	user := data.User{
		Name:     request.PostFormValue("name"),
		Email:    request.PostFormValue("email"),
		Password: request.PostFormValue("password"),
	}
	if err := user.Create(); err != nil {
		danger(err, "Cannot create user")
	}
	http.Redirect(writer, request, "/login", 302)
}

// POST /authenticate
// Authenticate the user given the email and password
func authenticate(writer http.ResponseWriter, request *http.Request) {
	// 리퀘스트 폼해석
	err := request.ParseForm()
	// email 키를 가진 폼 값을 가지고 사용자를 찾고
	user, err := data.UserByEmail(request.PostFormValue("email"))
	if err != nil {
		danger(err, "Cannot find user")
	}
	// 찾은 사용자의 암호와 폼에 입력된 암호가 같으면
	if user.Password == data.Encrypt(request.PostFormValue("password")) {
		// 세션을 만들고
		session, err := user.CreateSession()
		if err != nil {
			danger(err, "Cannot create session")
		}
		// 쿠키를 생성한 후
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		// 클라이언트에 생성한 쿠키를 보낸다.
		http.SetCookie(writer, &cookie)
		// 그리고는 루트경로로 리다이렉트
		http.Redirect(writer, request, "/", 302)
	// 암호가 일치하지 않으면
	} else {
		// 로그인 라우트로 리다이렉트
		http.Redirect(writer, request, "/login", 302)
	}

}

// GET /logout
// Logs the user out
func logout(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("_cookie")
	if err != http.ErrNoCookie {
		warning(err, "Failed to get cookie")
		session := data.Session{Uuid: cookie.Value}
		session.DeleteByUUID()
	}
	http.Redirect(writer, request, "/", 302)
}
