package main

import (
	"fmt"
)

// Tutorial of using the DIG DI >> https://blog.drewolson.org/go-dependency-injection-with-wire

// ## If any problem occurs try running these
// # At first make sure that $GOPATH/bin is in the PATH.
// If it is try running these: go get -u golang.org/x/tools/go/packages
// go get github.com/google/wire/cmd/wire

func main() {

	// // without DI
	// config := NewConfig()
	// personRespository := NewPersonRepository(config, &DB{})
	// personService := NewPersonService(config, personRespository)
	// personService.FindAll()
	// fmt.Println(NewConfig().DatabasePath)

	// with go-wire
	service, err := BuildDependencies()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	service.FindAll()
}
