package webserver

import (
	"github.com/Jvls1/go-ecommerce/product-service/domain"
	"github.com/Jvls1/go-ecommerce/product-service/test/mocks"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestCreateProduct(t *testing.T) {
	body := `{"name": "Product Test", "description": "This is a test product", "image_url": "https://example.com/image.jpg", "price": 19.99, "quantity": 10, "department_id": "880da008-32e4-4083-8b8a-43f9f3c9bbf2"}`
	bodyReader := strings.NewReader(body)

	req := httptest.NewRequest("POST", "/products", bodyReader)
	w := httptest.NewRecorder()

	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	productService := mocks.NewMockProductService(ctrl)
	product := domain.NewProduct("Product Test", "This is a test product", "https://example.com/image.jpg", 19.99, 10, "880da008-32e4-4083-8b8a-43f9f3c9bbf2")
	productService.EXPECT().CreateProduct(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(product, nil)
	productHandler := NewProductHandler(productService)
	productHandler.CreateProduct(c)

	assert.Equal(t, http.StatusCreated, c.Writer.Status())
	//TODO: Validate the response Body
}

func TestCreateProductInvalid(t *testing.T) {
	body := `{"name": "", "description": "This is a test product", "image_url": "", "price": -20, "quantity": -10, "department_id": ""}`
	bodyReader := strings.NewReader(body)

	req := httptest.NewRequest("POST", "/products", bodyReader)
	w := httptest.NewRecorder()

	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	productService := mocks.NewMockProductService(ctrl)

	productHandler := NewProductHandler(productService)
	productHandler.CreateProduct(c)

	assert.Equal(t, http.StatusBadRequest, c.Writer.Status())
}

func TestFindProducts(t *testing.T) {
	req := httptest.NewRequest("GET", "/products", nil)
	queryParams := req.URL.Query()
	queryParams.Add("page", "0")
	queryParams.Add("pageSize", "5")
	req.URL.RawQuery = queryParams.Encode()

	w := httptest.NewRecorder()

	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	productService := mocks.NewMockProductService(ctrl)
	productService.EXPECT().FindProducts(0, 5).Return([]*domain.Product{}, nil)

	productHandler := NewProductHandler(productService)
	productHandler.GetProducts(c)

	assert.Equal(t, http.StatusOK, c.Writer.Status())
}

func TestFindProductsInvalid(t *testing.T) {
	req := httptest.NewRequest("GET", "/products", nil)
	w := httptest.NewRecorder()

	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	productService := mocks.NewMockProductService(ctrl)

	productHandler := NewProductHandler(productService)
	productHandler.GetProducts(c)

	assert.Equal(t, http.StatusBadRequest, c.Writer.Status())
}

func TestFindProductById(t *testing.T) {
	productId := uuid.New().String()

	req := httptest.NewRequest("GET", "/product/"+productId, nil)
	w := httptest.NewRecorder()

	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = append(c.Params, gin.Param{Key: "id", Value: productId})

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	productService := mocks.NewMockProductService(ctrl)
	product := &domain.Product{
		ID:           productId,
		Name:         "Test",
		Description:  "Test",
		ImageURL:     "Test",
		Price:        20,
		Quantity:     1,
		DepartmentID: "1234",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	productService.EXPECT().FindProductByID(productId).Return(product, nil)
	productHandler := NewProductHandler(productService)
	productHandler.GetProductById(c)

	assert.Equal(t, http.StatusOK, c.Writer.Status())
}

func TestFindProductByIdInvalid(t *testing.T) {
	productId := "112234"

	req := httptest.NewRequest("GET", "/product/"+productId, nil)
	w := httptest.NewRecorder()

	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = append(c.Params, gin.Param{Key: "id", Value: productId})

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	productService := mocks.NewMockProductService(ctrl)

	productHandler := NewProductHandler(productService)
	productHandler.GetProductById(c)

	assert.Equal(t, http.StatusBadRequest, c.Writer.Status())
}

func TestFindProductByDepartmentId(t *testing.T) {
	departmentId := uuid.New().String()

	req := httptest.NewRequest("GET", "/product/department/"+departmentId, nil)
	w := httptest.NewRecorder()

	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = append(c.Params, gin.Param{Key: "departmentID", Value: departmentId})

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	productService := mocks.NewMockProductService(ctrl)

	productService.EXPECT().FindProductsByDepartmentId(departmentId).Return([]*domain.Product{}, nil)
	productHandler := NewProductHandler(productService)
	productHandler.GetProductByDepartmentID(c)

	assert.Equal(t, http.StatusOK, c.Writer.Status())
}

func TestFindProductByDepartmentIdInvalid(t *testing.T) {
	departmentId := "112234"

	req := httptest.NewRequest("GET", "/product/department/"+departmentId, nil)
	w := httptest.NewRecorder()

	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = append(c.Params, gin.Param{Key: "departmentID", Value: departmentId})

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	productService := mocks.NewMockProductService(ctrl)

	productHandler := NewProductHandler(productService)
	productHandler.GetProductByDepartmentID(c)

	assert.Equal(t, http.StatusBadRequest, c.Writer.Status())
}
