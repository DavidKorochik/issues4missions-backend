package handlers

import (
	"net/http"

	"github.com/DavidKorochik/issues4missions-backend/pkg/models"
	"github.com/DavidKorochik/issues4missions-backend/pkg/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h *Handler) CreateDepartment(c *gin.Context) {
	var req models.CreateDepartmentRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	newDepartment := models.Department{DepartmentName: req.DepartmentName}

	if err := h.departmentStore.CreateDepartment(newDepartment); err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, newDepartment)
}

func (h *Handler) GetDepartments(c *gin.Context) {
	departments, err := h.departmentStore.GetDepartments()

	if err != nil {

		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, util.ErrorResponse(err))
			return
		}

		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, departments)
}

func (h *Handler) GetDepartmentByName(c *gin.Context) {
	var uri models.GetDepartmentByNameURI

	if err := c.BindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	department, err := h.departmentStore.GetDepartmentByName(uri.DepartmentName)

	if err != nil {

		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, util.ErrorResponse(err))
			return
		}

		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, department)
}

func (h *Handler) UpdateDepartment(c *gin.Context) {
	var req models.UpdateDepartmentRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	departmentUpdates := models.Department{DepartmentName: req.DepartmentName}

	updatedDepartment, err := h.departmentStore.UpdateDepartment(departmentUpdates)

	if err != nil {

		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, util.ErrorResponse(err))
			return
		}

		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, updatedDepartment)
}

func (h *Handler) DeleteDepartment(c *gin.Context) {
	var uri models.DeleteDepartmentURI

	if err := c.BindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	deletedDepartment, err := h.departmentStore.DeleteDepartment(uri.DepartmentName)

	if err != nil {

		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, util.ErrorResponse(err))
			return
		}

		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, deletedDepartment)
}
