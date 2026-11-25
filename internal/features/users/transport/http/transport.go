package users_transport_http

type UsersHTTPHandler struct {
	usersService UserService
}

type UserService interface {
}

func NewUsersHTTPHandler(
	userService UserService,
) *UsersHTTPHandler {
	return &UsersHTTPHandler{
		usersService: userService,
	}
}
