package routes

func (r *Routes) AuthRoutes() {
	api := r.router.Group("/api")
	{
		api.POST("/auth", r.handler.AuthUser)
	}
}
