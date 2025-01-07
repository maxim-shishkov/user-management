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

func (get CreateRequest) Validate() error {
	return nil
}

type UpdateRequest struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
type UpdateResponse struct{}

func (get UpdateRequest) Validate() error {
	return nil
}

type DeleteRequest struct {
	ID int `json:"id"`
}
type DeleteResponse struct{}

func (get DeleteRequest) Validate() error {
	return nil
}

type GetRequest struct {
	ID int `json:"id"`
}
type GetResponse struct {
	User *User `json:"user"`
}

func (get GetRequest) Validate() error {
	return nil
}

type ListRequest struct{}
type ListResponse struct {
	Users []*User `json:"users"`
}

func (get ListRequest) Validate() error {
	return nil
}
