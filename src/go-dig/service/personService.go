package service

import (
	"fmt"
	"go-practice/src/go-dig/config"
	"go-practice/src/go-dig/model"
	"go-practice/src/go-dig/repository"
)

type PersonService struct {
	config            *config.Config
	personRespository *repository.PersonRepository
}

func (service *PersonService) FindAll() []*model.Person {
	if service.config.Enabled {
		fmt.Println("Config enabled")
	}
	fmt.Println("Service Called")
	return service.personRespository.FindAll()
}

func NewPersonService(config *config.Config, personRespository *repository.PersonRepository) *PersonService {
	return &PersonService{
		config:            config,
		personRespository: personRespository,
	}
}
