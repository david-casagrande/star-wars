package Planets

import (
  "github.com/david-casagrande/star-wars/request"
  "net/http"
  "encoding/json"
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
}

type PlanetJSON struct {
  Count int `json:"count"`
  Next string `json:"next"`
  Previous string `json:"previous"`
  Results []Planet `json:"results"`
}

type Planets struct {
  JSON []PlanetJSON
  Errors []error
}

func (p *Planets) populate() () {
  url := "http://swapi.co/api/planets"
  err := Request.Get(url, p.callback)

  if err != nil {
    return
  }

  remainingRequests := Request.RemainingRequests(p.JSON[0].Count)
  Request.All(url, remainingRequests, p.callback)
}

func (p *Planets) callback(resp *http.Response) () {
  data := PlanetJSON{}
  if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
    return
  }
  p.JSON = append(p.JSON, data)
}

func(p *Planets) all() ([]Planet) {
  results := []Planet{}

  for _, planet := range p.JSON {
    results = append(results, planet.Results...)
  }

  return results
}

func All() ([]Planet, error) {
  planets := Planets{}
  planets.populate()

  return planets.all(), nil
}
