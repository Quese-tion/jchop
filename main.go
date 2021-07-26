package main

import (
  "encoding/json"
  "fmt"
  "html/template"
  "log"
  "net/http"
  "os"
)

type User struct { //USER Data Model
  First string `json:"first_name"`
  Last string `json:"last_name"`
  Quote string
}
type Chuck struct { //USER Data Model
  Type string `json:"type"`
  Value Value `json:"value"`
}
type Value struct {
  Id int `json:"id"`
  Joke string `json:"joke"`
  Categories []string `json:"categories"`
}

func indexHandler(w http.ResponseWriter, r *http.Request){
  if r.URL.Path != "/" {
    http.NotFound(w, r)
    return}

  user:=new(User)
  userdata,err:=http.Get("https://names.mcquay.me/api/v0/")
  if err != nil{log.Fatal(err)}
  _ = json.NewDecoder(userdata.Body).Decode(&user)
  defer userdata.Body.Close()

  path := "./index.html"
  html, _ := template.ParseFiles(path)
  er := html.Execute(w, user)
  if er != nil {
    w.WriteHeader(http.StatusInternalServerError)
    log.Fatal(er)}}

func combineApi (w http.ResponseWriter, r *http.Request){
  if r.Method!=http.MethodGet || r.URL.Path != "/combine" {
    http.NotFound(w, r)
    return}

  r.ParseForm()
  var list =make(map[string]string)
   First := r.FormValue("First")
  list["First"]=First
  fmt.Println(First)
  chuck:=new(Chuck)
  chuckdata, er:=http.Get("https://api.icndb.com/jokes/random?firstName=John&lastName=Doe&limitTo=nerdy")
  if er != nil{log.Fatal(er)}
  _ = json.NewDecoder(chuckdata.Body).Decode(&chuck)
  defer chuckdata.Body.Close()

  ///Parse to work oage
    path:="./combine.html" //RIGHT NOW WE DO BIRTHDAY
    html, _ := template.ParseFiles(path)
    er = html.Execute(w,nil)
    if er != nil{
      w.WriteHeader(http.StatusInternalServerError)
      log.Fatal(er)
  }
}
func main()  {
  http.HandleFunc("/combine", combineApi)
  http.HandleFunc("/", indexHandler)
  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
    log.Printf("Defaulting to port %s", port)
  }
  log.Printf("Listening on port %s", port)
  log.Printf("Open http://localhost:%s in the browser", port)
  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))

}


