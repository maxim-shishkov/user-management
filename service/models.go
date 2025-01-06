package service

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type GetUserRequest struct {
	ID int `json:"id"`
}

type UpdateUserRequest struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type DeleteUserRequest struct {
	ID string `json:"id"`
}
