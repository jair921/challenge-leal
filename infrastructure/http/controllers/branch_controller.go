package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jair921/challenge-leal/internal/domain/entities"
	"github.com/jair921/challenge-leal/internal/domain/services"
	"net/http"
)

type BranchController struct {
	branchService services.BranchService
}

func NewBranchController(branchService services.BranchService) *BranchController {
	return &BranchController{branchService: branchService}
}

func (bc *BranchController) CreateBranch(c *gin.Context) {
	var branch entities.Branch
	if err := c.ShouldBindJSON(&branch); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := bc.branchService.CreateBranch(&branch); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, branch)
}

func (bc *BranchController) GetBranchByID(c *gin.Context) {
	id := c.Param("id")
	branch, err := bc.branchService.GetBranchByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, branch)
}

func (bc *BranchController) ListBranches(c *gin.Context) {
	branches, err := bc.branchService.ListBranches()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, branches)
}
