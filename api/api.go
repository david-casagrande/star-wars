package main

import (
  "net/http"
  "log"
  "github.com/julienschmidt/httprouter"
)

func main() {
  router := httprouter.New()

  router.HandleFunc("GET", "/api/:object", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
    log.Println(r)
    http.ServeFile(w, r, "../json/" + params.ByName("object") + ".json")
  })

  log.Fatal(http.ListenAndServe(":8080", router))
}
