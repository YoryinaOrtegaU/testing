package shark

import (
	"fmt"
	"math/rand"
	"testdoubles/prey"
	"testdoubles/simulator"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type whiteShark struct {
	// speed in m/s
	Speed float64
	// position of the shark in the map of 500 * 500 meters
	Position [2]float64
	// simulator
	Simulator simulator.CatchSimulator
}

func (w *whiteShark) Hunt(prey prey.Prey) error {
	if w.Simulator.CanCatch(w.Simulator.GetLinearDistance(w.Position), w.Speed, prey.GetSpeed()) {
		fmt.Println("ñam ñam")
		return nil
	}
	return fmt.Errorf("could not hunt the prey")
}

func CreateWhiteShark(simulator simulator.CatchSimulator) Shark {
	return &whiteShark{
		Speed:     144,
		Position:  GetInitialPosition(),
		Simulator: simulator,
	}
}

func GetInitialPosition() [2]float64 {
	x := rand.Float64() * 500
	rand.Seed(time.Now().UnixNano())
	y := rand.Float64() * 500
	return [2]float64{x, y}
}
