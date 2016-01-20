package People

import (
  "github.com/david-casagrande/star-wars/request"
  "github.com/david-casagrande/star-wars/utils"
  "net/http"
  "encoding/json"
  "log"
)

type Person struct {
  Name string `json:"name"`
  Height string `json:"height"`
  Mass string `json:"mass"`
  HairColor string `json:"hair_color"`
  SkinColor string `json:"skin_color"`
  Homeworld string `json:"homeworld"`
  Films []string `json:"films"`
  Species []string `json:"species"`
  Vehicles []string `json:"vehicles"`
  Starships []string `json:"starships"`
}

type PersonJSON struct {
  Count int `json:"count"`
  Next string `json:"next"`
  Previous string `json:"previous"`
  Results []Person `json:"results"`
}

type People struct {
  JSON []PersonJSON
  Errors []error
}

func(p *People) all() ([]Person) {
  results := []Person{}

  for _, person := range p.JSON {
    results = append(results, person.Results...)
  }

  return results
}

func (p *People) callback(resp *http.Response) () {
  data := PersonJSON{}
  if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
    log.Println(err)

    return
  }

  p.JSON = append(p.JSON, data)
}

func (p *People) populate() () {
  url := "http://swapi.co/api/people"
  err := Request.Get(url, p.callback)

  if err != nil {
    return
  }

  remainingRequests := Utils.RemainingRequests(p.JSON[0].Count, len(p.JSON[0].Results))
  Request.All(url, remainingRequests, p.callback)
}

func All() ([]Person, error) {
  people := People{}
  people.populate()

  return people.all(), nil
}
