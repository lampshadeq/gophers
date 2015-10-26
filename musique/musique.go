package musique

import (
  "html/template"
  "net/http"
  
  "appengine"
  "appengine/datastore"
  "appengine/user"
)

type Vote struct {
  User     string
  Musician string
}

type Data struct {
  Total  int
  MusMap map[string]int
}

func musiqueKey(c appengine.Context) *datastore.Key {
  return datastore.NewKey(c, "Musique", "default_musique", 0, nil)
}

func init() {
  http.HandleFunc("/", root)
  http.HandleFunc("/vote", vote)
}

func root(w http.ResponseWriter, r *http.Request) {
  // Get context for app engine
  c := appengine.NewContext(r)
  
  // Create query to access datastore
  q := datastore.NewQuery("Vote").Ancestor(musiqueKey(c))
  
  // Create and populate slice for data
  votes := make([]Vote, 0, 500)
  if _, err := q.GetAll(c, &votes); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  
  // Create and populate map to get counts
  m := make(map[string]int)
  t := 0
  for _, val := range votes {
    t++
    m[val.Musician]++
  }
  
  // Place total and map together
  d := Data {
    Total: t,
    MusMap: m,
  }

  // Launch the HTML template
  err := musTemp.Execute(w, d)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}

func vote(w http.ResponseWriter, r *http.Request) {
  // Get context for app engine
  c := appengine.NewContext(r)
  
  // Get the user
  u := user.Current(c)
  
  // If not logged in, force them
  if u == nil {
    url, err := user.LoginURL(c, r.URL.String())
    
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    
    w.Header().Set("Location", url)
    w.WriteHeader(http.StatusFound)
    return
  }
  
  // Ensure no blank entries
  if r.FormValue("choice") != "" {
    // Create a struct to hold info
    v := Vote {
      User: u.String(),
      Musician: r.FormValue("choice"),
    }
    
    // Create a key for the datastore
    key := datastore.NewIncompleteKey(c, "Vote", musiqueKey(c))
    
    // Push the info to the datastore
    _, err := datastore.Put(c, key, &v)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
  }
  
  // Redirect back to root
  http.Redirect(w, r, "/", http.StatusFound)
}

var musTemp = template.Must(template.New("mus").Parse(`
<!DOCTYPE html>
<html>
<head>
  <title>Musique: Visual Voting</title>
  <link rel="stylesheet" type="text/css" href="http://meyerweb.com/eric/tools/css/reset/reset.css">
  <link rel="stylesheet" type="text/css" href="https://fonts.googleapis.com/css?family=Permanent+Marker">
  
  <style>
    body {
      background-color: rgb(106,6,159);
      font-family: "Permanent Marker", cursive;
      padding-top: 3em;
    }
    
    h1 {
      font-size: 5em;
      width: 5em;
      margin: 0 auto;
    }
    
    h2 {
      font-size: 1.7em;
      width: 13em;
      margin: 0 auto;
    }
    
    form {
      width: 75%;
      margin: 0 auto;
    }
    
    h3 {
      font-size: 1.5em;
      margin-bottom: 0.5em;
    }
    
    #txt {
      background-color: rgb(106,6,159);
      width: 30%;
      border: 2px solid black;
      color: white;
      font-family: "Permanent Marker", cursive;
      padding: 0.5em 1em;
      font-size: 1.1em;
    }
    
    #txt:hover, #txt:focus {
      background-color: rgb(82,4,123);
    }
    
    #v {
      display: block;
      margin-top: 0.5em;
      font-family: "Permanent Marker", cursive;
      font-size: 1.1em;
      background-color: rgb(106,6,159);
      border: 2px solid black;
    }
    
    #v:hover {
      cursor: pointer;
      background-color: rgb(82,4,123);
      box-shadow: 0 0.2em 1em 0 rgb(157,113,232);
    }
    
    #main {
      width: 75%;
      margin: 0 auto;
      margin-top: 2em;
      border-top: 2px solid black;
      border-bottom: 2px solid black;
      padding-top: 1em;
      padding-bottom: 1em;
    }
    
    #main div {
      margin-bottom: 0.5em;
    }
    
    h4 {
      font-size: 3em;
      margin-bottom: 0.5em;
    }
    
    h5 {
      display: inline-block;
      width: 150px;
      font-size: 1.2em;
    }
    
    span {
      display: inline-block;
      height: 2em;
      width: 0;
      background-color: rgb(0,128,0);
    }
    
    h6 {
      width: 250px;
      margin: 0 auto;
      padding-top: 1em;
    }
  </style>
</head>
<body>
  <h1>Musique</h1>
  <h2>Who's your favorite?</h2>
  
  <form action="/vote" method="post">
    <h3>Enter your choice:</h3>
    <input id="txt" type="text" name="choice">
    <input id="v" type="submit" value="Vote!">
  </form>
  
  <div id="main">
    <h4>Results (alphabetical)</h4>
    {{range $key, $val := .MusMap}}
      <div>
        <h5>{{$key}}</h5>
        <span>{{$val}}</span>
      </div>
    {{end}}
  </div>
  
  <h6>Copyright &copy; 2015 - lampshadeq</h6>
  
  <script>
    // Get all the span elements
    var s = document.getElementsByTagName("SPAN");
    
    // Iterate
    for (var i = 0; i < s.length; ++i) {
      // Find the fraction (count / total)
      var f = parseInt(s[i].innerHTML) / {{.Total}};
      
      // Adjust width to fraction
      s[i].style.width = (f * 100).toString() + "%";
      
      // Delete the number
      s[i].innerHTML = ""
    }
  </script>
</body>
</html>
`))