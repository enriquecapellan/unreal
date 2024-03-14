package main

// import (
// 	"fmt"

// 	"github.com/go-playground/validator/v10"
// 	"github.com/kataras/iris/v12"
// )

// func main() {
// 	app := iris.New()
// 	app.Validator = validator.New()

// 	userRouter := app.Party("/user")
// 	{
// 		userRouter.Post("/", postUser)
// 	}
// 	app.Listen(":8080")
// }

// // User contains user information.
// type User struct {
// 	FirstName      string     `json:"fname" validate:"required"`
// 	LastName       string     `json:"lname" validate:"required"`
// 	Age            uint8      `json:"age" validate:"gte=0,lte=130"`
// 	Email          string     `json:"email" validate:"required,email"`
// 	FavouriteColor string     `json:"favColor" validate:"hexcolor|rgb|rgba"`
// 	Addresses      []*Address `json:"addresses" validate:"required,dive,required"`
// }

// // Address houses a users address information.
// type Address struct {
// 	Street string `json:"street" validate:"required"`
// 	City   string `json:"city" validate:"required"`
// 	Planet string `json:"planet" validate:"required"`
// 	Phone  string `json:"phone" validate:"required"`
// }

// type validationError struct {
// 	ActualTag string `json:"tag"`
// 	Namespace string `json:"namespace"`
// 	Kind      string `json:"kind"`
// 	Type      string `json:"type"`
// 	Value     string `json:"value"`
// 	Param     string `json:"param"`
// }

// func wrapValidationErrors(errs validator.ValidationErrors) []validationError {
// 	validationErrors := make([]validationError, 0, len(errs))
// 	for _, validationErr := range errs {
// 		validationErrors = append(validationErrors, validationError{
// 			ActualTag: validationErr.ActualTag(),
// 			Namespace: validationErr.Namespace(),
// 			Kind:      validationErr.Kind().String(),
// 			Type:      validationErr.Type().String(),
// 			Value:     fmt.Sprintf("%v", validationErr.Value()),
// 			Param:     validationErr.Param(),
// 		})
// 	}

// 	return validationErrors
// }

// func postUser(ctx iris.Context) {
// 	var user User
// 	err := ctx.ReadJSON(&user)
// 	if err != nil {
// 		// Handle the error, below you will find the right way to do that...

// 		if errs, ok := err.(validator.ValidationErrors); ok {
// 			// Wrap the errors with JSON format, the underline library returns the errors as interface.
// 			validationErrors := wrapValidationErrors(errs)

// 			// Fire an application/json+problem response and stop the handlers chain.
// 			ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
// 				Title("Validation error").
// 				Detail("One or more fields failed to be validated").
// 				Type("/user/validation-errors").
// 				Key("errors", validationErrors))

// 			return
// 		}

// 		// It's probably an internal JSON error, let's dont give more info here.
// 		ctx.StopWithStatus(iris.StatusInternalServerError)
// 		return
// 	}

// 	ctx.JSON(iris.Map{"message": "OK"})
// }
