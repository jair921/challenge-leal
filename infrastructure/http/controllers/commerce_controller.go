package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jair921/challenge-leal/internal/domain/entities"
	"github.com/jair921/challenge-leal/internal/domain/services"
	"net/http"
)

type CommerceController struct {
	commerceService services.CommerceService
}

func NewCommerceController(commerceService services.CommerceService) *CommerceController {
	return &CommerceController{commerceService: commerceService}
}

func (cc *CommerceController) CreateCommerce(c *gin.Context) {
	var commerce entities.Commerce
	if err := c.ShouldBindJSON(&commerce); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := cc.commerceService.CreateCommerce(&commerce); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, commerce)
}

func (cc *CommerceController) GetCommerceByID(c *gin.Context) {
	id := c.Param("id")
	commerce, err := cc.commerceService.GetCommerceByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, commerce)
}

func (cc *CommerceController) ListCommerces(c *gin.Context) {
	commerces, err := cc.commerceService.ListCommerces()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, commerces)
}
