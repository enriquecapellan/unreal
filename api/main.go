package main

import (
	"github.com/enriquecapellan/unreal/controllers"
	"github.com/enriquecapellan/unreal/initializers"
	"github.com/go-playground/validator/v10"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func init() {
	initializers.ConnectToDatabase()
	initializers.SyncDB()
}

func main() {
	app := iris.New()
	app.Validator = validator.New()

	cors := cors.New(cors.Options{
    AllowedOrigins:   []string{"*"},
    AllowCredentials: true,
	})

	app.UseRouter(cors)

	carsAPI := app.Party("/cars")
	{
		carsAPI.Use(iris.Compression)
		m := mvc.New(carsAPI)
		m.Handle(new(controllers.CarController))
	}

	app.Get("/", func(ctx iris.Context) {
		ctx.JSON(struct{ Message string }{Message: "Hello, World!"})
	})

	app.Listen(":8080")
}
