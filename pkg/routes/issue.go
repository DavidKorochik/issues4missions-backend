package routes

func (r *Routes) IssueRoutes() {
	api := r.router.Group("/api")

	api.POST("/issues", r.handler.CreateIssue)
	api.GET("/issues", r.handler.GetIssues)
	api.GET("/issues/:id", r.handler.GetIssueByID)
	api.PUT("/issues/:id", r.handler.UpdateIssue)
	api.DELETE("/issues/:id", r.handler.DeleteIssue)
}
