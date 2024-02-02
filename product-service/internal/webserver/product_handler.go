package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/Jvls1/go-ecommerce/product-service/domain"
	"github.com/Jvls1/go-ecommerce/product-service/internal/service"
)

type ProductHandler struct {
	ProductService *service.ProductService
}

func NewProductHandler(productService *service.ProductService) *ProductHandler {
	return &ProductHandler{ProductService: productService}
}

func (productHandler *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product domain.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := productHandler.ProductService.CreateProduct(product.Name, product.Description, product.ImageURL, product.Price, product.Quantity, product.DepartmentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)
}

func (productHandler *ProductHandler) GetProductById(w http.ResponseWriter, r *http.Request) {
}

func (productHandler *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
}

func (productHandler *ProductHandler) GetProductByDepartmentID(w http.ResponseWriter, r *http.Request) {
}

func (productHandler *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
}

func (productHandler *ProductHandler) SoftDeleteProduct(w http.ResponseWriter, r *http.Request) {
}
