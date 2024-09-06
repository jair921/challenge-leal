package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jair921/challenge-leal/internal/domain/entities"
	"github.com/jair921/challenge-leal/internal/domain/services"
	"net/http"
)

type CampaignController struct {
	campaignService services.CampaignService
}

func NewCampaignController(campaignService services.CampaignService) *CampaignController {
	return &CampaignController{campaignService: campaignService}
}

func (cc *CampaignController) CreateCampaign(c *gin.Context) {
	var campaign entities.Campaign
	if err := c.ShouldBindJSON(&campaign); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := cc.campaignService.CreateCampaign(&campaign); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, campaign)
}

func (cc *CampaignController) GetByID(c *gin.Context) {
	id := c.Param("id")
	campaign, err := cc.campaignService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, campaign)
}

func (cc *CampaignController) GetCampaignsByCommerceAndBranch(c *gin.Context) {
	commerceID := c.Param("commerceID")
	branchID := c.Param("branchID")
	campaigns, err := cc.campaignService.GetCampaignsByCommerceAndBranch(commerceID, branchID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, campaigns)
}

func (cc *CampaignController) GetActiveCampaigns(c *gin.Context) {
	commerceID := c.Param("commerceID")
	branchID := c.Param("branchID")
	campaigns, err := cc.campaignService.GetActiveCampaigns(commerceID, branchID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, campaigns)
}

func (cc *CampaignController) UpdateCampaign(c *gin.Context) {
	var campaign entities.Campaign
	if err := c.ShouldBindJSON(&campaign); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := cc.campaignService.UpdateCampaign(&campaign); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, campaign)
}

func (cc *CampaignController) DeleteCampaign(c *gin.Context) {
	id := c.Param("id")
	if err := cc.campaignService.DeleteCampaign(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
