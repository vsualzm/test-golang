package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"test-golang/helper"
	"test-golang/model"
	"test-golang/service"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.CreateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, product)
}

func (h *ProductHandler) GetProductByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	product, err := h.service.GetProductByID(uint(id))
	if err != nil {
		helper.RespondError(c, http.StatusNotFound, "Produk tidak ditemukan")
		return
	}

	helper.RespondJSON(c, http.StatusOK, "Data produk ditemukan", product)
}

func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	products, err := h.service.GetAllProducts()
	if err != nil {
		helper.RespondError(c, http.StatusInternalServerError, "Gagal mendapatkan data produk")
		return
	}

	fmt.Println("PRODUCT : ", products)

	helper.RespondJSON(c, http.StatusOK, "Data produk ditemukan", products)

}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.UpdateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := h.service.DeleteProduct(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}
