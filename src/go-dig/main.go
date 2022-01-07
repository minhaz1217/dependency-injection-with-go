package main

import (
	"go-practice/src/go-dig/config"
	"go-practice/src/go-dig/core"
	"go-practice/src/go-dig/repository"
	"go-practice/src/go-dig/service"
)

func main() {
	config := config.NewConfig()
	personRespository := repository.NewPersonRepository(config, &core.DB{})
	personService := service.NewPersonService(config, personRespository)
	personService.FindAll()

}
