package main

import (
  "html/template"
  "net/http"
)

var tpl *template.Template

func init() {
  http.HandleFunc("/", root)
  http.HandleFunc("/signin", signin)
  http.HandleFunc("/register", register)
  http.HandleFunc("/main", mainz)
  http.HandleFunc("/profile", profile)
  http.HandleFunc("/game", game)
  http.HandleFunc("/signout", signout)
  
  http.HandleFunc("/validate", validate)
  http.HandleFunc("/create", create)
  
  tpl = template.Must(template.ParseGlob("html/*.html"))
}

func root(w http.ResponseWriter, r *http.Request) {
  tpl.ExecuteTemplate(w, "index.html", nil)
}

func signin(w http.ResponseWriter, r *http.Request) {
  tpl.ExecuteTemplate(w, "signin.html", nil)
}

func register(w http.ResponseWriter, r *http.Request) {
  tpl.ExecuteTemplate(w, "register.html", nil)
}

func mainz(w http.ResponseWriter, r *http.Request) {
  tpl.ExecuteTemplate(w, "main.html", nil)
}

func profile(w http.ResponseWriter, r *http.Request) {
  tpl.ExecuteTemplate(w, "profile.html", nil)
}

func game(w http.ResponseWriter, r *http.Request) {
  tpl.ExecuteTemplate(w, "game.html", nil)
}

func signout(w http.ResponseWriter, r *http.Request) {
  http.Redirect(w, r, "/", 302)
}

func validate(w http.ResponseWriter, r *http.Request) {
  http.Redirect(w, r, "/main", 302)
}

func create(w http.ResponseWriter, r *http.Request) {
  http.Redirect(w, r, "/main", 302)
}