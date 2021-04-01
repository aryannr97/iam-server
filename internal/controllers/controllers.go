package controllers

// IControllers encapsulates all controller actions
type IControllers interface {
	IUserController
}

// IAMControllers encapsulates all the controllers
type IAMControllers struct {
	UserController IUserController
}
