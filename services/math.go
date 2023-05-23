package services

import (
	. "math/big"
)

type MathService struct {
}

func (ctrl MathService) Add(x *Float, y *Float) string {
	return x.Add(x, y).String()
}

func (ctrl MathService) Subtract(x *Float, y *Float) string {
	return x.Sub(x, y).String()
}

func (ctrl MathService) Multiply(x *Float, y *Float) string {
	return x.Mul(x, y).String()
}

func (ctrl MathService) Divide(x *Float, y *Float) string {
	return x.Quo(x, y).String()
}
