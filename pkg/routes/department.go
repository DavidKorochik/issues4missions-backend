package routes

import "github.com/DavidKorochik/issues4missions-backend/pkg/middleware"

func (r *Routes) DepartmentRoutes() {
	api := r.router.Group("/api").Use(middleware.AuthUser())
	{
		api.POST("/departments", r.handler.CreateDepartment)
		api.GET("/departments", r.handler.GetDepartments)
		api.GET("/departments/:department_name", r.handler.GetDepartmentByName)
		api.PUT("/departments/:department_name", r.handler.UpdateDepartment)
		api.DELETE("/departments/:department_name", r.handler.DeleteDepartment)
	}
}
