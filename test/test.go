package main

import (
  "log"
  "encoding/json"
  "github.com/david-casagrande/star-wars/objects/planets"
  "github.com/david-casagrande/star-wars/utils"

)

func main() {
  log.Println("started")
  planets, _ := Planets.All()

  data, err := json.Marshal(planets)
  if err != nil {
    log.Println(err)
    return
  }

  writeErr := Utils.WriteJSON("planets", data)
  if writeErr != nil {
    log.Println(writeErr)
    return
  }

  log.Println("finished")
}
