package main

import (
  "encoding/json"
  "testing"
)



func TestData(t *testing.T)  {
  chuckdata:= []byte(`{"ty": "success", "value": { "id": 534, "joke": "John Doe is the ultimate mutex, all threads fear him.", "categories": ["nerdy"] } }`)
  chuck:=new(Chuck)
  err:=json.Unmarshal(chuckdata, &chuck)
  if err !=nil {t.Fatal(err)}
}
