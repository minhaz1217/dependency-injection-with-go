package main

import (
	"fmt"
)

// Tutorial of using the DIG DI >> https://blog.drewolson.org/dependency-injection-in-go

func main() {

	// without DI
	config := NewConfig()
	personRespository := NewPersonRepository(config, &DB{})
	personService := NewPersonService(config, personRespository)
	personService.FindAll()
	fmt.Println(NewConfig().DatabasePath)
}
