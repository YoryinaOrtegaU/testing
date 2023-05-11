package shark

import (
	"errors"
	"testdoubles/prey"
)

type Shark interface {
	Hunt(prey prey.Prey) error
}

var (
	ErrNotCatch = errors.New("could not hunt the prey")
)
