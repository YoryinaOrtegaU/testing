package prey

type PreyStub struct {
	GetSpeedFn func() float64
}

func NewPreyStub() *PreyStub {
	return &PreyStub{}
}

func (ps *PreyStub) GetSpeed() float64 {
	return ps.GetSpeedFn()
}
