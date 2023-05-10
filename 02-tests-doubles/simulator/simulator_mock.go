package simulator

import "github.com/stretchr/testify/mock"

type CanSimulatorMockTestify struct {
	mock.Mock
}

// constructor
func NewCanSimulatorMockTestify() *CanSimulatorMockTestify {
	return &CanSimulatorMockTestify{}
}

func (c *CanSimulatorMockTestify) CanCatch(distance float64, speed float64, catchSpeed float64) bool {
	//Called le dice al objeto simulado que se ha llamado a un método y
	//obtiene una serie de argumentos para devolver.
	args := c.Called(distance, speed, catchSpeed)

	//obtiene el argumento en el índice especificado. retorna un bool
	indiceCero := args.Bool(0)

	return indiceCero
}

func (c *CanSimulatorMockTestify) GetLinearDistance(position [2]float64) float64 {
	args := c.Called(position)
	//Get Returns the argument at the specified index. retorna cualquier cosa
	indiceCero := args.Get(0).(float64)
	return indiceCero
}

// ------------------------------------------------------------
// manually
type CatchSimulatorMock struct {
	MethodCanCatch struct {
		Distance   float64
		Speed      float64
		CatchSpeed float64
		Ret0       bool
	}
	MethodGetLinearDistance struct {
		Position [2]float64
		Ret0     float64
	}
	// spy
	Spys map[string]bool
}

func (m *CatchSimulatorMock) CanCatch(distance float64, speed float64, catchSpeed float64) bool {
	// spy
	m.Spys["CanCatch"] = true

	m.MethodCanCatch.Distance = distance
	m.MethodCanCatch.Speed = speed
	m.MethodCanCatch.CatchSpeed = catchSpeed
	return m.MethodCanCatch.Ret0
}

func (m *CatchSimulatorMock) GetLinearDistance(position [2]float64) float64 {
	// spy
	m.Spys["GetLinearDistance"] = true

	m.MethodGetLinearDistance.Position = position
	return m.MethodGetLinearDistance.Ret0
}
