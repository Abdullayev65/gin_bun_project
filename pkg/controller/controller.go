package controller

type Controller struct {
	UserController
	PostController
	CommentController
}

func New() *Controller {
	return &Controller{}
}
