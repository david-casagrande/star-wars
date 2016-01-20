package Utils

import (
  "io/ioutil"
  "log"
  "math"
)

func WriteJSON(name string, data []byte) (error) {
  err := ioutil.WriteFile("../json/" + name + ".json", data, 0644)

  if err != nil {
    return err
  }

  return nil
}

func RemainingRequests(total int, resultsPerRequest int) (int) {
  if(total <= resultsPerRequest) {
    return 0
  }

  remainingTotal := total - resultsPerRequest
  roundedRequests := math.Ceil(float64(remainingTotal) / float64(resultsPerRequest))
  requests := int(roundedRequests)
  
  log.Println(requests)

  return int(requests)
}
