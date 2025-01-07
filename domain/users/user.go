package users

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
type CreateResponse struct {
	Id int `json:"id"`
}

type UpdateRequest struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
type UpdateResponse struct{}

type DeleteRequest struct {
	ID int `json:"id"`
}
type DeleteResponse struct{}

type GetRequest struct {
	ID int `json:"id"`
}
type GetResponse struct {
	User User `json:"user"`
}

type ListRequest struct{}
type ListResponse struct {
	Users []*User `json:"users"`
}
