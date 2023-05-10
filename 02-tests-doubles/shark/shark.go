package shark

import (
	"testdoubles/prey"
)

type Shark interface {
	Hunt(prey prey.Prey) error
}
