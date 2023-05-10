package shark

import (
	"testdoubles/prey"
	"testdoubles/simulator"
)

// sharkDummy implements the shark interface.
type SharkDummy struct {
	speed     float64
	Position  [2]float64
	simulator simulator.CatchSimulator
}

// constructor
func NewSharkDummy(simulator simulator.CatchSimulator) *SharkDummy {
	return &SharkDummy{
		speed:     144,
		Position:  [2]float64{3.6, 5.2},
		simulator: simulator,
	}
}

func (s *SharkDummy) Hunt(prey prey.Prey) error {
	return nil
}
