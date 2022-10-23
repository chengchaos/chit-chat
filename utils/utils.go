package utils

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	//"github.com/chengchaos/chit-chat/data"
	//"github.com/chengchaos/chit-chat/data"
)

type Configuration struct {
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
	Static       string
}

type Session struct {
	Id        int
	Uuid      string
	Email     string
	UserId    int
	CreatedAt time.Time
}

var Config *Configuration
var Logger *log.Logger

// P <br />
// Convenience function for printing to stdout
func P(a ...interface{}) {
	fmt.Println(a)
}

func initLog() {
	//file, err := os.OpenFile("chitchat.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//if err != nil {
	//	log.Fatalln("Failed to open log file", err)
	//}
	//Logger = log.New(file, "INFO", log.Ldate|log.Ltime|log.Lshortfile)
	Logger = log.New(os.Stdout, "INFO", log.Ldate|log.Ltime|log.Lshortfile)
}

func initConfig() {
	file, err := os.Open("Config.json")
	if err != nil {
		log.Fatalln("Cannot open Config file", err)
	}
	decoder := json.NewDecoder(file)
	Config = &Configuration{}
	err = decoder.Decode(&Config)
	LogInfo("Config", Config)
	if err != nil {
		log.Fatalln("Cannot get configuration from file", err)
	}
}

func ParseTemplateFiles(filenames ...string) (t *template.Template) {
	var files []string
	t = template.New("layout")
	for _, file := range filenames {
		fmt.Printf("file : => templates/%s.html \n", file)
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	t = template.Must(t.ParseFiles(files...))
	return
}

func GenerateHtml(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

// ErrorMessage Convenience function to redirect to the error message page
func ErrorMessage(w http.ResponseWriter, r *http.Request, msg string) {
	url := []string{"/err?msg=", msg}
	http.Redirect(w, r, strings.Join(url, ""), 302)
}

// for logging
func LogInfo(args ...interface{}) {
	Logger.SetPrefix("INFO ")
	Logger.Println(args...)
}

func LogError(args ...interface{}) {
	Logger.SetPrefix("ERROR ")
	Logger.Println(args...)
}

func LogWarning(args ...interface{}) {
	Logger.SetPrefix("WARNING ")
	Logger.Println(args...)
}

func Version() string {
	return "0.1"
}

func init() {
	initLog()
	initConfig()
}
