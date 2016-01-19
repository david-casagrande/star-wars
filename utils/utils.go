package Utils

import (
  "io/ioutil"
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
  requests := remainingTotal / resultsPerRequest

  if (remainingTotal % requests) > 0 {
    requests = requests + 1
  }

  return requests
}
