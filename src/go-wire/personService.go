package main

import (
	"fmt"
)

type PersonService struct {
	config            *Config
	personRespository *PersonRepository
}

func (service *PersonService) FindAll() []*Person {
	if service.config.Enabled {
		fmt.Println("Config enabled")
	}
	fmt.Println("Service Called")
	return service.personRespository.FindAll()
}

func NewPersonService(config *Config, personRespository *PersonRepository) *PersonService {
	return &PersonService{
		config:            config,
		personRespository: personRespository,
	}
}
