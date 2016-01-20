package Request

import (
  "net/http"
  // "encoding/json"
  "log"
  "strconv"
)

type Data interface {
}

type callback func(*http.Response) ()

// type Data struct {
//   Count int `json:"count"`
//   Next string `json:"next"`
//   Previous string `json:"previous"`
//   Results []interface{} `json:"results"`
// }

// func paginatedAsync(requests int, url string, data Data) ([]Data, []error) {
//   successes := make(chan Data, requests)
//   errors := make(chan error, requests)
//
//   for i := 0; i < requests; i++ {
//     go func(i int) {
//       pagedUrl := url + "?page=" + strconv.Itoa(i + 2)
//       d, err := Get(pagedUrl, &data)
//       if err != nil {
//         errors <- err
//       } else {
//         successes <- d
//       }
//     }(i)
//   }
//
//   allData := []Data{}
//   allErrors := []error{}
//
//   for i := 0; i < requests; i++ {
//     select {
//     case data := <-successes:
//       allData = append(allData, data)
//     case err := <-errors:
//       allErrors = append(allErrors, err)
//     }
//   }
//
//   return allData, allErrors
//   // return Planets{ JSON: allPlanets }, allErrors
// }

func Get(url string, fn callback) (error) {
  resp, err := http.Get(url)
  if err != nil {
    return err
  }

  defer resp.Body.Close()

  fn(resp)

  return nil
}

func All(url string, requests int, fn callback) {
  errors := make(chan error)

  for i := 0; i < requests; i++ {
    go func(i int) {
      pagedUrl := url + "?page=" + strconv.Itoa(i + 2)
      log.Println(pagedUrl)
      errors <- Get(pagedUrl, fn)
    }(i)
  }

  for i := 0; i < requests; i++ {
    select {
    case err := <-errors:
      if err != nil {
        log.Println(err)
      }
    }
  }
}
