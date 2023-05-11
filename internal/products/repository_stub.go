package products

// RepositoryStub is and struct and contains the method get all by seller
type RepositoryStub struct {
	GetAllBySellerStub func(sellerID string) ([]Product, error)
}

// NewRepositoryStub is a builder and creates a new implementation of RepositoryStub
func NewRepositoryStub() *RepositoryStub {
	return &RepositoryStub{}
}

// GetAllBySeller predefined implementation of mock behavior for the GetAllBySeller method of the actual Repository object
func (rs *RepositoryStub) GetAllBySeller(sellerID string) ([]Product, error) {
	return rs.GetAllBySellerStub(sellerID)
}