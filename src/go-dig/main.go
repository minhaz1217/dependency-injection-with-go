package main

// Tutorial of using the DIG DI >> https://blog.drewolson.org/dependency-injection-in-go
import (
	"fmt"
	"go-practice/src/go-dig/config"
	"go-practice/src/go-dig/core"
	"go-practice/src/go-dig/repository"
	"go-practice/src/go-dig/service"

	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	container := dig.New()
	err := container.Provide(config.NewConfig)
	if err != nil {
		fmt.Println(1, err)
		panic(err)
	}
	err = container.Provide(core.NewDB)
	if err != nil {
		fmt.Println(2, err)
		panic(err)
	}
	err = container.Provide(repository.NewPersonRepository)
	if err != nil {
		fmt.Println(3, err)
		panic(err)
	}
	err = container.Provide(service.NewPersonService)
	if err != nil {
		fmt.Println(4, err)
		panic(err)
	}
	fmt.Println("Container built")
	return container
}

func main() {

	// without DI
	// config := config.NewConfig()
	// personRespository := repository.NewPersonRepository(config, &core.DB{})
	// personService := service.NewPersonService(config, personRespository)
	// personService.FindAll()

	// with DI
	container := BuildContainer()
	err := container.Invoke(func(service *service.PersonService) {
		service.FindAll()
	})

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}
