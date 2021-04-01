package services

// IServices encapsulates all services actions
type IServices interface {
	IUserService
}

// IAMServices encapsulates all the services
type IAMServices struct {
	UserService IUserService
}
