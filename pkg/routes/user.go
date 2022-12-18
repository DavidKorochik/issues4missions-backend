package routes

func (r *Routes) UserRoutes() {
	api := r.router.Group("/api")
	{
		api.POST("/users", r.handler.CreateUser)
		api.GET("/users", r.handler.GetUsers)
		api.GET("/users/:id", r.handler.GetUserByID)
		api.PUT("/users/:id", r.handler.UpdateUser)
		api.DELETE("/users/:id", r.handler.DeleteUser)
	}
}
