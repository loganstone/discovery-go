package funcs

import "math"

// Func is return float
type Func func(float64) float64

// Transform is return function
type Transform func(Func) Func

const tolerance = 0.00001
const dx = 0.00001

// Square is return multiply by itself
func Square(x float64) float64 {
	return x * x
}

// FixedPoint is some
func FixedPoint(f Func, firstGuess float64) float64 {
	closeEnough := func(v1, v2 float64) bool {
		return math.Abs(v1-v2) < tolerance
	}
	var try Func
	try = func(guess float64) float64 {
		next := f(guess)
		if closeEnough(guess, next) {
			return next
		}
		return try(next)
	}
	return try(firstGuess)
}

// FixedPointOfTeansform is return transform
func FixedPointOfTeansform(g Func, transform Transform, guess float64) float64 {
	return FixedPoint(transform(g), guess)
}

// Deriv is return differential
func Deriv(g Func) Func {
	return func(x float64) float64 {
		return (g(x+dx) - g(x)) / dx
	}
}

// NewtonTransform is do Newton Method
func NewtonTransform(g Func) Func {
	return func(x float64) float64 {
		return x - (g(x) / Deriv(g)(x))
	}
}

// Sqrt is return a square root
func Sqrt(x float64) float64 {
	return FixedPointOfTeansform(func(y float64) float64 {
		return Square(y) - x
	}, NewtonTransform, 1.0)
}
