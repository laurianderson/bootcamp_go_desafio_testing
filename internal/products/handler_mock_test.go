package products

import (
	"net/http"
	"testing"
	"github.com/bootcamp-go/desafio-cierre-testing/pkg"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHandlerMock_GetSellers(t *testing.T) {
	//succed case
	t.Run("when seller id is", func(t *testing.T) {
		// arrange
		sellerIDToSearch := "FEX112AC"
		expectedSeller := []Product{
			{
				ID:          "mock",
				SellerID:    "FEX112AC",
				Description: "generic product",
				Price:       123.55,
			},
		}
		service := NewServiceMock()
		handler := NewHandler(service)

		// act
		service.On("GetAllBySeller", sellerIDToSearch).Return(expectedSeller, nil)

		// create server
		r := gin.New()
		r.GET("/api/v1/products" , handler.GetProducts)
		req, resp := pkg.CreateRequestTest(http.MethodGet, "/api/v1/products?seller_id=" + sellerIDToSearch , "")
		r.ServeHTTP(resp, req)

		// assert
		assert.Equal(t, http.StatusOK, resp.Code)
		service.AssertExpectations(t)
	})

	


}
