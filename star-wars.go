package main

import (
  "log"
  "net/http"
  "encoding/json"
  "strconv"
  "io/ioutil"
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

func remainingRequests(firstPage PlanetJSON, url string) (Planets, error) {
  allPlanets := []PlanetJSON{ firstPage, }
  maxResults := 10
  remainingTotal := firstPage.Count - maxResults
  requests := remainingTotal / maxResults

  if (remainingTotal % requests) > 0 {
    requests = requests + 1
  }

  for i := 0; i < requests; i++ {
    pagedUrl := url + "?page=" + strconv.Itoa(i + 2)
    data, err := request(pagedUrl)
    if err != nil {
      return Planets{}, err
    }
    allPlanets = append(allPlanets, data)
  }

  return Planets{ JSON: allPlanets }, nil
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

func main() {
  url := "http://swapi.co/api/planets"
  data, err := request(url)

  if err != nil {
    log.Println(err)
  }

  planets, err := remainingRequests(data, url)

  writeJSON(planets.All())
  log.Println(planets.All())
}
