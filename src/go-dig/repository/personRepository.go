package repository

import (
	"fmt"
	"go-practice/src/go-dig/config"
	"go-practice/src/go-dig/core"
	"go-practice/src/go-dig/model"
)

type PersonRepository struct {
	config   *config.Config
	database *core.DB
}

func (repo *PersonRepository) FindAll() []*model.Person {
	people := []*model.Person{}
	for i := 0; i < 5; i++ {
		people = append(people, &model.Person{})
	}
	fmt.Println("Repo called")
	return people
}

func NewPersonRepository(config *config.Config, database *core.DB) *PersonRepository {
	return &PersonRepository{
		config:   config,
		database: database,
	}
}
