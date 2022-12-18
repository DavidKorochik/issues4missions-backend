package routes

import (
	"github.com/DavidKorochik/issues4missions-backend/pkg/middleware"
)

func (r *Routes) IssueRoutes() {
	api := r.router.Group("/api").Use(middleware.AuthUser())
	{
		api.POST("/issues", r.handler.CreateIssue)
		api.GET("/issues", r.handler.GetIssues)
		api.GET("/issues/:id", r.handler.GetIssueByID)
		api.GET("/issues/department", r.handler.GetIssueByDepartment)
		api.PUT("/issues/:id", r.handler.UpdateIssue)
		api.DELETE("/issues/:id", r.handler.DeleteIssue)
	}
}
