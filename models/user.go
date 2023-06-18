package models

type UserPrimaryKey struct {
	Id string `json:"id"`
}

type CreateUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Balance   int    `json:"balance"`
}

type User struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Balance   int    `json:"balance"`
}

type UpdateUser struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Balance   int    `json:"balance"`
}

type UserGetListRequest struct {
	Offset int
	Limit  int
}

type UserGetListResponse struct {
	Count int
	Users []*User
}
