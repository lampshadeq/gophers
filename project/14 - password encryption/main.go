package main

import (
  "html/template"
  "net/http"
  "encoding/json"
  "time"
  
  "appengine"
  "appengine/datastore"
  "appengine/memcache"
  "golang.org/x/crypto/bcrypt"
  
  "github.com/nu7hatch/gouuid"
)

var tpl *template.Template
var fuq bool

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
  fuq = false
}

func root(w http.ResponseWriter, r *http.Request) {
  fuq = false
  tpl.ExecuteTemplate(w, "index.html", nil)
}

func signin(w http.ResponseWriter, r *http.Request) {
  tpl.ExecuteTemplate(w, "signin.html", fuq)
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
  c := appengine.NewContext(r)
  o, _ := r.Cookie("session")
  sd := memcache.Item {
    Key:        o.Value,
    Value:      []byte(""),
    Expiration: time.Duration(1 * time.Microsecond),
  }
  memcache.Set(c, &sd)
  o.MaxAge = -1
  http.SetCookie(w, o)
  http.Redirect(w, r, "/", 302)
}

func validate(w http.ResponseWriter, r *http.Request) {
  c := appengine.NewContext(r)
  k := datastore.NewKey(c, "Users", r.FormValue("username"), 0, nil)
  var u User
  err := datastore.Get(c, k, &u)
  if err != nil ||
     bcrypt.CompareHashAndPassword([]byte(u.Password),
     []byte(r.FormValue("password"))) != nil {
    fuq = true
    http.Redirect(w, r, "/signin", 302)
  } else {
    fuq = false
    createSession(w, r, u)
    http.Redirect(w, r, "/main", 302)
  }
}

func create(w http.ResponseWriter, r *http.Request) {
  c := appengine.NewContext(r)
  h, _ := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), bcrypt.DefaultCost)
  u := User {
    Username: r.FormValue("username"),
    Password: string(h),
    Email:    r.FormValue("email"),
  }
  k := datastore.NewKey(c, "Users", u.Username, 0, nil)
  _, err := datastore.Put(c, k, &u)
  if err != nil {
    http.Redirect(w, r, "/register", 302)
  }
  createSession(w, r, u)
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

func createSession(w http.ResponseWriter, r *http.Request, u User) {
  c := appengine.NewContext(r)
  
  id, _ := uuid.NewV4()
  o := &http.Cookie {
    Name:   "session",
    Value:  id.String(),
    Path:   "/",
    //Secure: true,
    //HttpOnly: true,
  }
  http.SetCookie(w, o)
  
  j, _ := json.Marshal(u)
  sd := memcache.Item {
    Key:        id.String(),
    Value:      j,
    Expiration: time.Duration(20 * time.Second),
  }
  memcache.Set(c, &sd)
}