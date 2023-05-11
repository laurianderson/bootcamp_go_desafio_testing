package products

import "github.com/stretchr/testify/mock"

// RepositoryMock is and struct and contains an embeding package mock
type RepositoryMock struct {
	mock.Mock
}

// NewRepositoryMock is a builder and creates a new implementation of RepositoryMock
func NewRepositoryMock() *RepositoryMock {
	return &RepositoryMock{}
}

// GetAllBySeller predefined implementation of mock behavior for the GetAllBySeller method of the actual Repository object
func (repoMock *RepositoryMock) GetAllBySeller(sellerID string) ([]Product, error) {
	// input (expectations)
	// called I get the arguments passed by the parameter of the function
	args := repoMock.Called(sellerID)

	//output
	//args.Get(0) get the first argument as a []Product, which is assigned to the seller variable
	seller := args.Get(0).([]Product)
	// args.Error(1) get the second argument as an error, which is assigned to the variable err
	err := args.Error(1)

	return seller, err
}
