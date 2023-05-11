package sharktest

import (
	"testdoubles/prey"
	"testdoubles/shark"
	"testdoubles/simulator"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHunt(t *testing.T) {
	//Arrage
	presa := prey.NewPreyStub()
	presa.GetSpeedFn = func() (speedPrey float64) {
		speedPrey = 5
		return
	}

	simu := simulator.NewCanSimulatorMockTestify()

	//Act
	simu.On("GetLinearDistance", mock.Anything).Return(2.4)
	simu.On("CanCatch", 2.4, 144.0, 5.0).Return(true)

	sharneido := shark.CreateWhiteShark(simu)
	err := sharneido.Hunt(presa)

	//Assert
	assert.Empty(t, err)
	//AssertExpectations afirma que todo lo especificado con On y Return
	//se llamó de hecho como se esperaba. Las llamadas pueden haber ocurrido en cualquier orden.
	simu.AssertExpectations(t)
}
