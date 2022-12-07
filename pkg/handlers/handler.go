package handlers

type HandlerRepo struct {
	userStore  user.Store
	issueStore issue.Store
}

func NewHandlerRepo(us user.Store, is issue.Store) *HandlerRepo {
	return &HandlerRepo{
		userStore:  us,
		issueStore: is,
	}
}
