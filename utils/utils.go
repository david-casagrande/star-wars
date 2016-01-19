package Utils

import (
  "io/ioutil"
  "log"
)

func WriteJSON(name string, data []byte) (error) {
  // var data interface{}
  // data, err := json.Marshal(planets)
  // if err != nil {
  //   log.Println(err)
  //   return
  // }

  err := ioutil.WriteFile("../json/" + name + ".json", data, 0644)

  if err != nil {
    log.Println(err)
  }

  return nil
}
