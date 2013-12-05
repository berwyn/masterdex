package controller

import (
	model "../model"
	"code.google.com/p/gorest"
	"github.com/eaigner/hood"
)

type SpeciesService struct {
	// Register the service root
	gorest.RestService `root:"/pokemon/" consumes:"application/json" produces:"application/json"`

	// Provide the database connection
	Database *hood.Hood

	// Register the endpoints
	pokemonDetails gorest.EndPoint `method:"GET" path:"/{ID:string}" output:"model.Species"`
}

func(serv SpeciesService) PokemonDetails(ID string) (m model.Species){
	var results []model.Species
	err := serv.Database.Where("dex_number", "=", ID).Limit(1).Find(&results)
	if err != nil {
		m = results[0]
		return
	} 
	serv.ResponseBuilder().SetResponseCode(404).Overide(true)
	return
}
