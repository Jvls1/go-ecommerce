package webserver

import (
	"net/http"
	"strconv"

	"github.com/Jvls1/go-ecommerce/product-service/domain"
	"github.com/Jvls1/go-ecommerce/product-service/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProductHandler struct {
	ProductService *service.ProductService
}

func NewProductHandler(productService *service.ProductService) *ProductHandler {
	return &ProductHandler{ProductService: productService}
}

func (productHandler *ProductHandler) CreateProduct(c *gin.Context) {
	var product domain.Product

	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := productHandler.ProductService.CreateProduct(product.Name, product.Description, product.ImageURL, product.Price, product.Quantity, product.DepartmentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result.ID)
}

func (productHandler *ProductHandler) GetProducts(c *gin.Context) {
	pageStr := c.Query("page")
	pageSizeStr := c.Query("pageSize")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		pageSize = 10
	}

	// Agora você tem os valores de page e pageSize como inteiros
	// Faça o que for necessário com esses valores, por exemplo, passá-los para o repository

	// Exemplo de resposta simples (substitua isso pelo seu código real)
	c.JSON(http.StatusOK, gin.H{"page": page, "pageSize": pageSize})
}

func (productHandler *ProductHandler) GetProductById(c *gin.Context) {
	productID := c.Param("id")
	err := uuid.Validate(productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	product, err := productHandler.ProductService.FindProductByID(productID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": product})
}

func (productHandler *ProductHandler) GetProductByDepartmentID(c *gin.Context) {
	departmentID := c.Param("departmentID")
	err := uuid.Validate(departmentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	products, err := productHandler.ProductService.FindProductsByDepartmentId(departmentID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found for this "})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": products})
}
