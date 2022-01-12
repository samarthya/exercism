package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.
const (
	Silly = "silly nephew, there cannot be %d cows"
)

var NegativeFodder = errors.New("negative fodder")
var ZeroDivision = errors.New("division by zero")

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	cows int
}

// error builtin interface implements the `Error() string` method
func (s SillyNephewError) Error() string {
	return fmt.Sprintf(Silly, s.cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	v, err := weightFodder.FodderAmount()

	if cows > 0 {
		if err == nil {
			switch {
			case  v > 0:
				return v/float64(cows), nil
			case v < 0 :
				return 0.0, NegativeFodder
			}
		}

		switch{
		case err == ErrScaleMalfunction:
			if v >0 {
				return (v*2)/float64(cows), nil
			}
			if v < 0 {
				return 0.0, NegativeFodder
			}
		}

		return 0.0, err
	} else if cows < 0 {
		return 0.0, SillyNephewError{cows}
	} 
		return 0.0, ZeroDivision
	
}
