package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

// hello world
// template
// struct
// iteration

//model
type TodoList struct {
	Title string
	Items []Todo
}

//model
type Todo struct {
	Job string
	Done bool

}

type UserMsg struct {
	username string
	message string
}

//controller
func index(w http.ResponseWriter, r *http.Request){
	data := TodoList{
		Title: "Todo List of Today",
		Items: []Todo{
			{Job:"Go to school",Done: true},
			{Job:"Do homework",Done: false},
		},
	}
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w,data)
}

//controller
func form(w http.ResponseWriter, r *http.Request){
	tmpl := template.Must(template.ParseFiles("form.html"))
	if r.Method == http.MethodGet{
		tmpl.Execute(w,nil)
		return
	}
	userMsg := UserMsg{
		username: r.FormValue("username"),
		message: r.FormValue("message"),
	}
	fmt.Println(userMsg.username,userMsg.message)
	tmpl.Execute(w, struct {
		Success bool
	}{true})

}

//Middleware
func logging(f http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		start := time.Now()
		log.Println(r.URL.Path,r.Method,time.Since(start))
		f(w,r)
	}
}

func main(){
	//address
	http.HandleFunc("/", logging(index))
	http.HandleFunc("/form",logging(form))
	// use static file
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/",http.StripPrefix("/static/",fs))
	log.Fatal(http.ListenAndServe(":5000",nil))
}
