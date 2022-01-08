package main

import (
	"github.com/google/wire"
)

func BuildDependencies() (*PersonService, error) {
	wire.Build(
		NewDB,
		NewConfig,
		NewPersonRepository,
		NewPersonService,
	)
	return &PersonService{}, nil
}
