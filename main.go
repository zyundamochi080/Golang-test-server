package main

import (
  "github.com/gorilla/mux"
  "net/http"
  "fmt"
  "encoding/json"

  "google.golang.org/appengine"
)

type Response struct {
    Name  string `json:"name"`
}

func main() {
  r := mux.NewRouter()
  r.HandleFunc("/api/v1/{dates}", UserHandler).Methods("GET")
  r.Host("localhost")
  http.Handle("/", r)
  appengine.Main()
}

type Reply struct {
  Message string `json:"message"`
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  dates := vars["dates"]

  reply := Reply {
    Message: "Date:" + dates,
  }

  json, _ := json.Marshal(reply)

  fmt.Fprint(w, string(json))
}
