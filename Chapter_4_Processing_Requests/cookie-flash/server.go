package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
)

// 쿠키 설정
func setMessage(w http.ResponseWriter, r *http.Request) {
	msg := []byte("Hello World!")
	c := http.Cookie{
		Name:  "flash",
		// 인코딩하는 이유는 특수문자 때문이다. 기본적으로 처리가 들어가있어야함
		Value: base64.URLEncoding.EncodeToString(msg),
	}
	http.SetCookie(w, &c)
}

// 쿠키 있으면 보여주고 MaxAge(ie 6,7 지원 안함) Expires (지원) 설정으로 삭제
func showMessage(w http.ResponseWriter, r *http.Request) {
	// 있던 쿠키 가져온 다음
	c, err := r.Cookie("flash")
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Fprintln(w, "No message found")
		}
	} else {
		// 기존 쿠기 기간 -1 로 해서
		rc := http.Cookie{
			Name:    "flash",
			MaxAge:  -1,
			Expires: time.Unix(1, 0),
		}
		// 덮어씌워서 클라에서 쿠키 삭제되게 하고
		http.SetCookie(w, &rc)
		// 앞서 가져왔던 값을
		val, _ := base64.URLEncoding.DecodeString(c.Value)
		// 응답에 표시하기
		fmt.Fprintln(w, string(val))
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/set_message", setMessage)
	http.HandleFunc("/show_message", showMessage)
	server.ListenAndServe()
}
