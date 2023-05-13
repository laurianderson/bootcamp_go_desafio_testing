package products

import "github.com/stretchr/testify/mock"


// NewServiceMock is a builder and creates a new implementation of ServiceMock
func NewServiceMock() *ServiceMock {
	return &ServiceMock{}
}


// ServiceMock is and struct and contains an embeding package mock
type ServiceMock struct {
	mock.Mock
}


// GetAllBySeller calls to repository mock and get all by seller
func (s *ServiceMock) GetAllBySeller(sellerID string) ([]Product, error) {
	args := s.Called(sellerID)
	err := args.Error(1)
	return args.Get(0).([]Product), err
}
