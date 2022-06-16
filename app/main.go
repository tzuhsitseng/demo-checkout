package main

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"

	"github.com/tzuhsitseng/demo-checkout/calculation"
	"github.com/tzuhsitseng/demo-checkout/configs"
)

func main() {
	if err := configs.Init(); err != nil {
		log.Panicln("cannot init configs:", err)
	}

	app := iris.New()
	app.Validator = validator.New()
	app.Use(iris.Compression)

	service := calculation.NewService()
	controller := calculation.NewController(service)

	app.Get("/calculation", controller.Calculate)
	_ = app.Listen(":8080")
}
