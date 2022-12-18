package routes

func (r *Routes) SessionRoutes() {
	api := r.router.Group("/api")
	{
		api.POST("/renew", r.handler.RenewAccessToken)
	}
}
