package main

import (
	"errors"
	"fmt"
	"./data"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"encoding/json"
)

// 설정 구조체 만들기
type Configuration struct {
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
	Static       string
}

// Configuration 타입 config 변수 정의, 구조체는 필드의 모음이라고 함. 클래스가 아님
// https://tour.golang.org/moretypes/2
var config Configuration

// log.Logger 구조체를 가리키는 포인터 리시버를 변수에 할당한다.
// 그냥 변수를 넣어주면 구조체가 아니라 변수 자체를 변경하려 하기 때문이다.
// https://go-tour-kr.appspot.com/#52
var logger *log.Logger

// Convenience function for printing to stdout
func p(a ...interface{}) {
	fmt.Println(a)
}


// 실행 순서 - 패키지 init() main.go 의 init()
// https://www.popit.kr/%EC%9E%90%EB%B0%94-%EA%B0%9C%EB%B0%9C%EC%9E%90%EA%B0%80-go-%EC%9E%A0%EA%B9%90-%EC%82%AC%EC%9A%A9%ED%95%B4-%EB%B4%A4%EC%8A%B5%EB%8B%88%EB%8B%A42/
func init() {
	// 설정 불러오기 함수 실행
	loadConfig()
	file, err := os.OpenFile("chitchat.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	logger = log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
}

// 설정 파일 불러오기
func loadConfig() {
	// 워킹디렉토리에서 설정 파일을 읽어서
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatalln("Cannot open config file", err)
	}
	// 그걸 json 디코더에 연결하고
	decoder := json.NewDecoder(file)

	// 구조체 리터럴로 빈 구조체 인스턴스 생성해서 변수에 할당
	config = Configuration{}
	err = decoder.Decode(&config) // 디코더로 디코드한 결과를 config 에 넣는다.
	// 포인터를 넘겨 줘야지 구조체 내부의 필드에 접근할 수 있다. 그래서 &으로 포인터 넘겨줌
	// https://tour.golang.org/moretypes/4
	if err != nil {
		log.Fatalln("Cannot get configuration from file", err)
	}
}

// Convenience function to redirect to the error message page
func error_message(writer http.ResponseWriter, request *http.Request, msg string) {
	url := []string{"/err?msg=", msg}
	http.Redirect(writer, request, strings.Join(url, ""), 302)
}

// Checks if the user is logged in and has a session, if not err is not nil
func session(writer http.ResponseWriter, request *http.Request) (sess data.Session, err error) {
	// 사용자 요청에서 "_cookei" 이름의 쿠키 정보를 가져오고
	cookie, err := request.Cookie("_cookie")
	if err == nil {
		// data 패키지 Session 구조체리터럴로 UUid 값만 설정해서 sess 에 할당
		sess = data.Session{Uuid: cookie.Value}
		// Session 구조체 리시버를 가진 함수 Check() 을 통해서 세션 확인
		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}

// parse HTML templates
// pass in a list of file names, and get a template
func parseTemplateFiles(filenames ...string) (t *template.Template) {
	var files []string
	t = template.New("layout")
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	t = template.Must(t.ParseFiles(files...))
	return
}

// 빈 인터페이스는 모든 것을 받을 수 있다. filenames 는 variadic 가변인 것으로 ... 를 통해 나열한다. es6 비구조화 같네
func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	// 문자열 배열 선언
	var files []string
	// 마치 for... in 처럼 filenames 를 range 를 통해 filenames 를 순회하면서 반복문 처리
	for _, file := range filenames {
		// files 에다가 현재 파일 이름을 가지고 만든 템플릿 파일 경로문자열을 누적
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}

	// 해당 템플릿파일 경로를 해석한 결과를 해석한 결과를 Must 로 얻는다.
	templates := template.Must(template.ParseFiles(files...))
	// 그럼 이제 템플릿 중에 layout 과 data 를 넘겨서 최종 html 을 writer 에 작성한다.
	templates.ExecuteTemplate(writer, "layout", data)
}

// for logging
func info(args ...interface{}) {
	logger.SetPrefix("INFO ")
	logger.Println(args...)
}

func danger(args ...interface{}) {
	logger.SetPrefix("ERROR ")
	logger.Println(args...)
}

func warning(args ...interface{}) {
	logger.SetPrefix("WARNING ")
	logger.Println(args...)
}

// version
func version() string {
	return "0.1"
}
