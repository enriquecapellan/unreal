package controllers

import (
	"github.com/enriquecapellan/unreal/initializers"
	"github.com/enriquecapellan/unreal/models"
	"github.com/enriquecapellan/unreal/utils"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
)

type CarController struct {
}

func (c *CarController) Get(ctx iris.Context) {
	cars := []models.Car{}
	initializers.DB.Find(&cars)
	ctx.JSON(cars)
}

func (c *CarController) GetBy(ctx iris.Context, ID int) {
	var car = models.Car{}
	initializers.DB.Find(&car, ID)
	ctx.JSON(car)
}

func (c *CarController) Post(ctx iris.Context) {
	var car models.Car
	err := ctx.ReadJSON(&car)

	if err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			// Wrap the errors with JSON format, the underline library returns the errors as interface.
			validationErrors := utils.WrapValidationErrors(errs)

			// Fire an application/json+problem response and stop the handlers chain.
			ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
				Title("Validation error").
				Detail("One or more fields failed to be validated").
				Type("/user/validation-errors").
				Key("errors", validationErrors))

			return
		}

		// It's probably an internal JSON error, let's dont give more info here.
		ctx.StopWithStatus(iris.StatusInternalServerError)
		return
	}

	// result := initializers.DB.Create(&car)

	// if result.Error != nil {
	// 	ctx.StatusCode(iris.StatusBadRequest)
	// 	return
	// }

	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(car)
}

func (c *CarController) DeleteBy(ctx iris.Context, ID int) {
	initializers.DB.Delete(&models.Car{}, ID)
	ctx.StatusCode(iris.StatusNoContent)
}
