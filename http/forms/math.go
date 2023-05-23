package forms

type InputParams struct {
	X string `validate:"required,numeric"`
	Y string `validate:"required,numeric"`
}
