package products

import "github.com/stretchr/testify/mock"


// NewServiceMock is a builder and creates a new implementation of ServiceMock
func NewServiceMock(repoMock *RepositoryMock) *ServiceMock {
	return &ServiceMock{
		repositoryMock: *repoMock,
	}
}


// ServiceMock is and struct and contains an embeding package mock
type ServiceMock struct {
	mock.Mock
	repositoryMock RepositoryMock
}


// GetAllBySeller calls to repository mock and get all by seller
func (s *ServiceMock) GetAllBySeller(sellerID string) ([]Product, error) {
	return s.repositoryMock.GetAllBySeller(sellerID)
}
