package main

import (
  "html/template"
  "net/http"
  
  "appengine"
  "appengine/datastore"
)

var tpl *template.Template

type User struct {
  Username  string
  Password  string
  Email     string
}

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
  http.HandleFunc("/checkUsername", checkUsername)
  
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
  c := appengine.NewContext(r)
  u := User {
    Username: r.FormValue("username"),
    Password: r.FormValue("password"),
    Email:    r.FormValue("email"),
  }
  k := datastore.NewKey(c, "Users", u.Username, 0, nil)
  _, err := datastore.Put(c, k, &u)
  if err != nil {
    http.Redirect(w, r, "/register", 302)
  }
  http.Redirect(w, r, "/main", 302)
}

func checkUsername(w http.ResponseWriter, r *http.Request) {
  c := appengine.NewContext(r)
  var u User
  k := datastore.NewKey(c, "Users", r.FormValue("username"), 0, nil)
  err := datastore.Get(c, k, &u)
  if err != nil {
    w.Write([]byte("good"))
  } else {
    w.Write([]byte("bad"))
  }
}