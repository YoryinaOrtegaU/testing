package prey

import "math/rand"

type tuna struct {
	// max speed in m/s
	maxSpeed  float64
}

func (t *tuna) GetSpeed() float64 {
	return t.maxSpeed * rand.Float64()
}

func CreateTuna() Prey {
	return &tuna{
		maxSpeed:  252,
	}
}
