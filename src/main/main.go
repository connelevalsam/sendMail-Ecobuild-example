package main

import (
	"fmt"
	"github.com/gopkg.in/gomail.v2"
	"html/template"
	"log"
	"net/http"
)

var (
	templ *template.Template
	err   error
)

func init() {
	templ, err = templ.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/contact", contact)
	http.Handle("/assets/", http.FileServer(http.Dir(".")))
	http.Handle("favicon.ico", http.NotFoundHandler())
	fmt.Println("running server on port :9000")
	http.ListenAndServe(":9000", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	err = templ.ExecuteTemplate(res, "index.html", nil)
}

func contact(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		s := req.FormValue("subject")
		e := req.FormValue("mail")
		m := req.FormValue("message")

		cb := gomail.NewMessage()
		cb.SetHeader("From", e)
		cb.SetHeader("To", "raillblaze@gmail.com")
		cb.SetHeader("Subject", s)
		cb.SetBody("text/plain", m)

		d := gomail.NewDialer("smtp.gmail.com", 587, "connelblaze@gmail.com", "elemietta")
		if err := d.DialAndSend(cb); err != nil {
			panic(err)
		}
		fmt.Println("message delivered successfully!")
	}

	err = templ.ExecuteTemplate(res, "contact.html", nil)
	if err != nil {
		log.Println(err)
	}
}
