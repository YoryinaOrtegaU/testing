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
	presa.GetSpeedFn = func() (speedPrey float64) {
		speedPrey = 5
		return
	}

	simu := simulator.NewCanSimulatorMockTestify()
	expectedPosition := [2]float64{3.6, 5.2}

	//Act
	shar := shark.NewSharkDummy(simu)
	simu.On("GetLinearDistance", expectedPosition).Return(2.4)
	simu.On("CanCatch", 2.4, 144, 5).Return(true)

	err := shar.Hunt(presa)

	/* sharneido := shark.CreateWhiteShark(simu)
	err := sharneido.Hunt(presa) */

	//Assert
	assert.Empty(t, err)
	//AssertExpectations afirma que todo lo especificado con On y Return
	//se llamó de hecho como se esperaba. Las llamadas pueden haber ocurrido en cualquier orden.
	simu.AssertExpectations(t)
}
