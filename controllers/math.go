package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	. "github.com/vlbaluk/apalon-test/http"
	"github.com/vlbaluk/apalon-test/http/forms"
	"github.com/vlbaluk/apalon-test/services/interfaces"

	"math/big"
	"net/http"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type MathController struct {
}

func (ctrl MathController) Calc(c *gin.Context, fn interfaces.Calculate, ch interfaces.CachedCalculate, operation string) {
	x := c.Query("x")
	y := c.Query("y")

	if ctrl.ValidateParams(x, y, c) != nil {
		return

	}

	biX, biY := big.Float{}, big.Float{}
	biX.SetString(x)
	biY.SetString(y)

	answer, cachedBool := ch(&biX, &biY, operation, fn)

	c.JSON(http.StatusOK, &Response{Action: operation, X: x, Y: y, Answer: answer, Cashed: cachedBool})
	return
}

func (ctrl MathController) ValidateParams(x string, y string, c *gin.Context) error {
	input := &forms.InputParams{X: x, Y: y}
	err := validate.Struct(input)

	if err != nil {
		errMessage := ""
		for _, error := range err.(validator.ValidationErrors) {
			if len(errMessage) > 0 {
				errMessage += ", "
			}
			if error.Tag() == "numeric" {
				errMessage += error.Field() + " parameter should be numeric"
			}
			if error.Tag() == "required" {
				errMessage += error.Field() + " parameter doesn't exist"
			}
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, &ErrorResponse{Message: errMessage})
		return err
	}
	return nil
}
