package sharktest

import (
	"testdoubles/prey"
	"testdoubles/shark"
	"testdoubles/simulator"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHunt(t *testing.T) {
	//Arrage
	presa := prey.NewPreyStub()
	simu := simulator.NewCanSimulatorMockTestify()

	//Act
	shar := shark.CreateWhiteShark(simu)
	err := shar.Hunt(presa)

	//Assert
	assert.NotEmpty(t, err)
	assert.Error(t, err, "could not hunt the prey")
}
