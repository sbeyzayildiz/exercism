package thefarm

import (
	"errors"
	"fmt"
)

type InvalidCowsError struct {
	deadCows int
	message  string
}

// Error returns a string when the silly nephew counts the dead cows with bovine spongiform encephalopathy
func (haveACow InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", haveACow.deadCows, haveACow.message)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder FodderCalculator, cows int) (float64, error) {
	// this is MY happy path
	if fodder, err := weightFodder.FodderAmount(cows); err == nil {
		if fatFactor, err := weightFodder.FatteningFactor(); err == nil {
			return fodder * fatFactor / float64(cows), nil
		} else {
			return 0, err
		}
	} else {
		return 0, err
	}
}
func ValidateInputAndDivideFood(weightFodder FodderCalculator, cows int) (float64, error) {
	if cows <= 0 {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(weightFodder, cows)
}
func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		return &InvalidCowsError{cows, "there are no negative cows"}
	}
	if cows == 0 {
		return &InvalidCowsError{cows, "no cows don't need food"}
	}
	return nil
}
