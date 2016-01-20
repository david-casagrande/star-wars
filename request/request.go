package Request

import (
  "net/http"
  "log"
  "strconv"
)

type Data interface {
}

type callback func(*http.Response) ()

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
