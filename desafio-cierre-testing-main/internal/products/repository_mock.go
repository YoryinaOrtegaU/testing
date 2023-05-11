package products

import "github.com/stretchr/testify/mock"

type RepositoryMockTestify struct {
	mock.Mock
}

func NewRepositoryMockTestify() *RepositoryMockTestify {
	return &RepositoryMockTestify{}
}

func (rp *RepositoryMockTestify) GetAllBySeller(sellerID string) (prodList []Product, err error) {

	args := rp.Called(sellerID)

	// output
	prodList = args.Get(0).([]Product)
	err = args.Error(1)
	return
}
