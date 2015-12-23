package main

import (
  "html/template"
  "net/http"
  "encoding/json"
  "time"
  "math/rand"
  "strconv"
  
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
  Wins      int
  Ties      int
  Losses    int
  Total     int
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
  http.HandleFunc("/getmove", getmove)
  
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
  i := getSession(r)
  if len(i.Value) > 0 {
    var u User
    json.Unmarshal(i.Value, &u)
    tpl.ExecuteTemplate(w, "profile.html", u)
  } else {
    tpl.ExecuteTemplate(w, "profile.html", nil)
  }
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
    Wins:     0,
    Ties:     0,
    Losses:   0,
    Total:    0,
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

func getSession(r *http.Request) *memcache.Item {
  o, err := r.Cookie("session")
  if err != nil {
    return &memcache.Item{}
  }
  c := appengine.NewContext(r)
  i, err := memcache.Get(c, o.Value)
  if err != nil {
    return &memcache.Item{}
  }
  return i
}

func getmove(w http.ResponseWriter, r *http.Request) {
  var d [9]string
  d[0] = r.FormValue("sq0")
  d[1] = r.FormValue("sq1")
  d[2] = r.FormValue("sq2")
  d[3] = r.FormValue("sq3")
  d[4] = r.FormValue("sq4")
  d[5] = r.FormValue("sq5")
  d[6] = r.FormValue("sq6")
  d[7] = r.FormValue("sq7")
  d[8] = r.FormValue("sq8")
  
  var m = getSession(r)
  var u User
  json.Unmarshal(m.Value, &u)
  
  // Case 0 (Top Horiz)
  if d[0] == d[1] && d[1] == d[2] {
    if d[0] == "u" {
      u.Wins += 1
      u.Total += 1
      w.Write([]byte("win"))
      return
    } else if d[0] == "c" {
      u.Losses += 1
      u.Total += 1
      w.Write([]byte("lose"))
      return
    }
  }
  
  // Case 1 (Mid Horiz)
  if d[3] == d[4] && d[4] == d[5] {
    if d[3] == "u" {
      u.Wins += 1
      u.Total += 1
      w.Write([]byte("win"))
      return
    } else if d[3] == "c" {
      u.Losses += 1
      u.Total += 1
      w.Write([]byte("lose"))
      return
    }
  }
  
  // Case 2 (Bot Horiz)
  if d[6] == d[7] && d[7] == d[8] {
    if d[6] == "u" {
      u.Wins += 1
      u.Total += 1
      w.Write([]byte("win"))
      return
    } else if d[6] == "c" {
      u.Losses += 1
      u.Total += 1
      w.Write([]byte("lose"))
      return
    }
  }
  
  // Case 4 (Left Vert)
  if d[0] == d[3] && d[3] == d[6] {
    if d[0] == "u" {
      u.Wins += 1
      u.Total += 1
      w.Write([]byte("win"))
      return
    } else if d[0] == "c" {
      u.Losses += 1
      u.Total += 1
      w.Write([]byte("lose"))
      return
    }
  }
  
  // Case 5 (Mid Vert)
  if d[1] == d[4] && d[4] == d[7] {
    if d[1] == "u" {
      u.Wins += 1
      u.Total += 1
      w.Write([]byte("win"))
      return
    } else if d[1] == "c" {
      u.Losses += 1
      u.Total += 1
      w.Write([]byte("lose"))
      return
    }
  }
  
  // Case 6 (Right Vert)
  if d[2] == d[5] && d[5] == d[8] {
    if d[2] == "u" {
      u.Wins += 1
      u.Total += 1
      w.Write([]byte("win"))
      return
    } else if d[2] == "c" {
      u.Losses += 1
      u.Total += 1
      w.Write([]byte("lose"))
      return
    }
  }
  
  // Case 7 (NWSE Cross)
  if d[0] == d[4] && d[4] == d[8] {
    if d[0] == "u" {
      u.Wins += 1
      u.Total += 1
      w.Write([]byte("win"))
      return
    } else if d[0] == "c" {
      u.Losses += 1
      u.Total += 1
      w.Write([]byte("lose"))
      return
    }
  }
  
  // Case 8 (NESW Cross)
  if d[2] == d[4] && d[4] == d[6] {
    if d[2] == "u" {
      u.Wins += 1
      u.Total += 1
      w.Write([]byte("win"))
      return
    } else if d[2] == "c" {
      u.Losses += 1
      u.Total += 1
      w.Write([]byte("lose"))
      return
    }
  }
  
  // Case 9 (Tie)
  c := 0
  for _, t := range d {
    if t != "n" {
      c++
    }
  }
  if c == 9 {
    u.Ties += 1
    u.Total += 1
    w.Write([]byte(strconv.Itoa(-1)))
    return
  }
  
  rand.Seed(time.Now().UTC().UnixNano())
  var i = 0
  var s = ""
  for s != "n" {
    i = rand.Intn(len(d))
    s = d[i]
  }
  
  w.Write([]byte(strconv.Itoa(i)))
}