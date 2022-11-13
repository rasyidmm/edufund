package entity

type RegisterRequest struct {
	Fullname string
	Username string
	Password string
}

type RegisterResponse struct {
	StatusCode string
	StatusDesc string
}
type UserRequest struct {
	Username string
}
type UserResponse struct {
	Fullname string
	Username string
	Password string
}
