package main

import (
  "log"
  "net/http"
  "encoding/json"
  "strconv"
  "io/ioutil"
  // "github.com/julienschmidt/httprouter"
)

type Planet struct {
  Name string `json:"name"`
  RotationPeriod string `json:"rotation_period"`
  OrbitalPeriod string `json:"orbital_period"`
  Diameter string `json:"diameter"`
  Climate string `json:"climate"`
  Gravity string `json:"gravity"`
  Terrain string `json:"terrain"`
  SurfaceWater string `json:"surface_water"`
  Population string `json:"population"`
      // "residents": [
      //     "http://swapi.co/api/people/5/",
      //     "http://swapi.co/api/people/68/",
      //     "http://swapi.co/api/people/81/"
      // ],
      // "films": [
      //     "http://swapi.co/api/films/6/",
      //     "http://swapi.co/api/films/1/"
      // ],
      // "created": "2014-12-10T11:35:48.479000Z",
      // "edited": "2014-12-20T20:58:18.420000Z",
      // "url": "http://swapi.co/api/planets/2/"
}

type PlanetJSON struct {
  Count int `json:"count"`
  Next string `json:"next"`
  Previous string `json:"previous"`
  Results []Planet `json:"results"`
}

type Planets struct {
  JSON []PlanetJSON
}

func (p Planets) All() ([]Planet) {
  planets := []Planet{}
  for _, planetJSON := range p.JSON {
    planets = append(planets, planetJSON.Results...)
  }

  return planets
}

func request(url string) (PlanetJSON, error) {
  resp, err := http.Get(url)
  if err != nil {
    return PlanetJSON{}, err
  }

  defer resp.Body.Close()

  var data PlanetJSON
  if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
    return PlanetJSON{}, err
  }

  return data, nil
}

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

func getPlanets() (Planets, error) {
  url := "http://swapi.co/api/planets"
  data, err := request(url)

  if err != nil {
    log.Println(err)
  }

  planets, _ := remainingRequests(data, url)
  // if err != nil {
  //   return Planets{}, err
  // }

  return planets, nil
}

func main() {
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

  planets, _ := getPlanets()
  writeJSON(planets.All())
}
