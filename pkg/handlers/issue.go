package handlers

import (
	"fmt"
	"net/http"

	"github.com/DavidKorochik/issues4missions-backend/pkg/models"
	"github.com/DavidKorochik/issues4missions-backend/pkg/token"
	"github.com/DavidKorochik/issues4missions-backend/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (h *Handler) CreateIssue(c *gin.Context) {
	var req models.CreateIssueRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)
	newIssue := models.Issue{Title: req.Title, Description: req.Description, UserID: authPayload.UserID}

	if err := h.issueStore.CreateIssue(newIssue); err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, newIssue)
}

func (h *Handler) GetIssues(c *gin.Context) {
	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)
	issues, err := h.issueStore.GetIssues(authPayload.UserID)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, util.ErrorResponse(err))
			return
		}

		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, issues)
}

func (h *Handler) GetIssueByID(c *gin.Context) {
	var uri models.GetIssueByIDUri

	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	issue, err := h.issueStore.GetIssueByID(uri.ID)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, util.ErrorResponse(err))
			return
		}

		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, issue)
}

func (h *Handler) GetIssueByDepartment(c *gin.Context) {
	var query models.GetIssueByDepartmentQuery

	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	issue, err := h.issueStore.GetIssueByDepartment(query.DepartmentName)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, util.ErrorResponse(err))
			return
		}

		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, issue)
}

func (h *Handler) UpdateIssue(c *gin.Context) {
	var uri models.UpdateIssueUri
	var req models.UpdateIssueRequest

	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)

	if !h.isUserManipulatesHisIssue(uri.ID, authPayload.UserID) {
		c.JSON(http.StatusUnauthorized, fmt.Errorf("User is not authorized"))
		return
	}

	updatedIssue, err := h.issueStore.UpdateIssue(uri.ID, req)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, util.ErrorResponse(err))
			return
		}

		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, updatedIssue)
}

func (h *Handler) DeleteIssue(c *gin.Context) {
	var uri models.DeleteIssueUri

	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)

	if !h.isUserManipulatesHisIssue(uri.ID, authPayload.UserID) {
		c.JSON(http.StatusUnauthorized, fmt.Errorf("User is not authorized"))
		return
	}

	deletedIssue, err := h.issueStore.DeleteIssue(uri.ID)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, util.ErrorResponse(err))
			return
		}

		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, deletedIssue)
}

func (h *Handler) isUserManipulatesHisIssue(issueID uuid.UUID, userID uuid.UUID) bool {
	issue, _ := h.issueStore.GetIssueByID(issueID)

	if issue.UserID != userID {
		return false
	}

	return true
}
