package main

import (
  "log"
  "encoding/json"
  "github.com/david-casagrande/star-wars/objects/planets"
  "github.com/david-casagrande/star-wars/objects/people"
  "github.com/david-casagrande/star-wars/utils"

)

func planets() {
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
}

func people() {
  people, _ := People.All()

  data, err := json.Marshal(people)
  if err != nil {
    log.Println(err)
    return
  }

  writeErr := Utils.WriteJSON("people", data)
  if writeErr != nil {
    log.Println(writeErr)
    return
  }
}

func main() {
  log.Println("started")
  // planets()
  people()
  log.Println("finished")
}
