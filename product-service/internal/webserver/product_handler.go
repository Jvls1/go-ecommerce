package webserver

import (
	"net/http"
	"strconv"

	"github.com/Jvls1/go-ecommerce/product-service/domain"
	"github.com/Jvls1/go-ecommerce/product-service/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func NewProductHandler(productService service.ProductService) ProductHandler {
	return &productHandler{productService}
}

type productHandler struct {
	ProductService service.ProductService
}

type ProductHandler interface {
	CreateProduct(c *gin.Context)
	GetProducts(c *gin.Context)
	GetProductById(c *gin.Context)
	GetProductByDepartmentID(c *gin.Context)
}

func (productHandler *productHandler) CreateProduct(c *gin.Context) {
	var product domain.Product

	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdProduct, err := productHandler.ProductService.CreateProduct(product.Name, product.Description, product.ImageURL, product.Price, product.Quantity, product.DepartmentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdProduct.ID)
}

func (productHandler *productHandler) GetProducts(c *gin.Context) {
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

	products, err := productHandler.ProductService.FindProducts(page, pageSize)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"items": products, "page": page, "pageSize": pageSize})
}

func (productHandler *productHandler) GetProductById(c *gin.Context) {
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

func (productHandler *productHandler) GetProductByDepartmentID(c *gin.Context) {
	departmentID := c.Param("departmentID")
	err := uuid.Validate(departmentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	products, err := productHandler.ProductService.FindProductsByDepartmentId(departmentID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found for this department"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": products})
}
