package controller

import (
  model "../model/species.go"
  "code.google.com/p/gorest"
)

type SpeciesService struct {
  // Register the service root
  gorest.RestService  `root:"/pokemon/" consumes:"application/json" produces:"application/json"`

  // Provide the database connection
  database        *hood.Hood

  // Register the endpoints
  pokemonDetails  gorest.EndPoint `method:"GET" path:"/{ID:int}" output:"model.Species"`
}

func (serv SpeciesService) PokemonDetails(ID int) model.Species {
  m, err := database.Where("dex_number", "=", ID).Limit(1).Find(&model.Species)
  if err != nil {
    return m
  } else {
    return Species{}
  }
}