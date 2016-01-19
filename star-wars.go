package main

import (
  "log"
  "net/http"
  "encoding/json"
  "strconv"
  "io/ioutil"
  // "github.com/julienschmidt/httprouter"
  "github.com/david-casagrande/star-wars/objects/planets"
)

// type Planet struct {
//   Name string `json:"name"`
//   RotationPeriod string `json:"rotation_period"`
//   OrbitalPeriod string `json:"orbital_period"`
//   Diameter string `json:"diameter"`
//   Climate string `json:"climate"`
//   Gravity string `json:"gravity"`
//   Terrain string `json:"terrain"`
//   SurfaceWater string `json:"surface_water"`
//   Population string `json:"population"`
// }

// type PlanetJSON struct {
//   Count int `json:"count"`
//   Next string `json:"next"`
//   Previous string `json:"previous"`
//   Results []Planet `json:"results"`
// }

// type Parser struct {
//   JSON []Data
// }

// func (p Parser) All() ([]Data) {
//   results := []Data{}
//   for _, j := range p.JSON {
//     results = append(results, j.Results...)
//   }
//
//   return results
// }

// func request(url string) (PlanetJSON, error) {
//   resp, err := http.Get(url)
//   if err != nil {
//     return PlanetJSON{}, err
//   }
//
//   defer resp.Body.Close()
//
//   var data PlanetJSON
//   if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
//     return PlanetJSON{}, err
//   }
//
//   return data, nil
// }
//
// type Data interface {
// }

// func Get(url string, data Data) (Data, error) {
//   resp, err := http.Get(url)
//   if err != nil {
//     return data, err
//   }
//
//   defer resp.Body.Close()
//
//   if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
//     return data, err
//   }
//
//   return data, nil
// }

func remainingRequests(firstPage PlanetJSON, url string) (Planets, []error) {
  maxResults := 10
  remainingTotal := firstPage.Count - maxResults
  requests := remainingTotal / maxResults

  if (remainingTotal % requests) > 0 {
    requests = requests + 1
  }

  // return remainingRequestsSync(firstPage, requests, url)
  return remainingRequestsAsync(firstPage, requests, url)
}

func remainingRequestsAsync(firstPage PlanetJSON, requests int, url string) (Planets, []error) {
  successes := make(chan PlanetJSON, requests)
  errors := make(chan error, requests)

  for i := 0; i < requests; i++ {
    go func(i int) {
      pagedUrl := url + "?page=" + strconv.Itoa(i + 2)
      data, err := request(pagedUrl)
      if err != nil {
        errors <- err
      } else {
        successes <- data
      }
    }(i)
  }

  allPlanets := []PlanetJSON{ firstPage, }
  allErrors := []error{}

  for i := 0; i < requests; i++ {
    select {
    case data := <-successes:
      log.Println(data)
      allPlanets = append(allPlanets, data)
    case err := <-errors:
      allErrors = append(allErrors, err)
    }
  }

  return Planets{ JSON: allPlanets }, allErrors
}

func remainingRequestsSync(firstPage PlanetJSON, requests int, url string) (Planets, []error) {
  allPlanets := []PlanetJSON{ firstPage, }
  errors := []error{}

  for i := 0; i < requests; i++ {
    pagedUrl := url + "?page=" + strconv.Itoa(i + 2)
    data, err := request(pagedUrl)
    if err != nil {
      errors = append(errors, err)
    } else {
      allPlanets = append(allPlanets, data)
    }
  }

  return Planets{ JSON: allPlanets }, errors
}

func writeJSON(planets []Planet) {
  // var data interface{}
  data, err := json.Marshal(planets)
  if err != nil {
    log.Println(err)
    return
  }

  fileError := ioutil.WriteFile("./json/planets.json", data, 0644)

  if fileError != nil {
    log.Println(fileError)
  }
}

func getPlanets() (Parser, error) {
  url := "http://swapi.co/api/planets"
  data, err := Get(url, &PlanetJSON{})//request(url)

  if err != nil {
    log.Println(err)
  }

  return Parser{ JSON: data, }

  // planets, _ := remainingRequests(data, url)
  // if err != nil {
  //   return Planets{}, err
  // }

  // return planets, nil
}

func main() {
  planets.All()
  // router := httprouter.New()
  //
  // router.GET("/planets", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
  //   planets, _ := getPlanets()
  //
  //   w.Header().Set("Content-Type", "application/json; charset=utf-8")
  //   json.NewEncoder(w).Encode(planets.All())
  // })
  //
  //
  // log.Fatal(http.ListenAndServe(":8080", router))

  // planets, _ := getPlanets()
  // writeJSON(planets.All())
}
