package main

import (
  "encoding/json"
  "fmt"
  "log"
  "net/http"
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

func main()  {
  user:=new(User)
  userdata,err:=http.Get("https://names.mcquay.me/api/v0/")
  if err != nil{log.Fatal(err)}
  _ = json.NewDecoder(userdata.Body).Decode(&user)
  defer userdata.Body.Close()

  chuck:=new(Chuck)
  chuckdata, er:=http.Get("http://api.icndb.com/jokes/random?firstName=John&lastName=Doe&limitTo=nerdy")
  if er != nil{log.Fatal(err)}
  _ = json.NewDecoder(chuckdata.Body).Decode(&chuck)
  defer chuckdata.Body.Close()

  fmt.Println("Hi ", user.First, ",\n\n Chuck Norris once said:\n\t\"", chuck, "\"")
}


