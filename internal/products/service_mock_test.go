package products

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestServiceMock_GetSellers(t *testing.T) {
	//succed case
	t.Run("get sellers", func(t *testing.T) {
		//arrange
		storage := NewRepositoryMock()
		storage.On("GetAllBySeller", "FEX112AC").Return([]Product{
			{ID: "mock", SellerID: "FEX112AC", Description: "generic product", Price: 123.55},
		}, nil)

		sv := NewService(storage)
		//act
		seller, err := sv.GetAllBySeller("FEX112AC")

		//assert
		assert.NoError(t, err)
		assert.Equal(t, []Product{
			{ID: "mock", SellerID: "FEX112AC", Description: "generic product", Price: 123.55},
		}, seller)
		storage.AssertExpectations(t)
	})
}
