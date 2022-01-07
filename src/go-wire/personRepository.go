package main

import (
	"fmt"
)

type PersonRepository struct {
	config   *Config
	database *DB
}

func (repo *PersonRepository) FindAll() []*Person {
	people := []*Person{}
	for i := 0; i < 5; i++ {
		people = append(people, &Person{})
	}
	fmt.Println("Repo called")
	return people
}

func NewPersonRepository(config *Config, database *DB) *PersonRepository {
	return &PersonRepository{
		config:   config,
		database: database,
	}
}
