package model

type Role struct {
	Role_Id int64  `json:"role_id"`
	Role    string `json:"role"`
}

type Permession_assignment struct {
	Permission_Id int64 `json:"permission_id"`
	Role_Id       int64 `json:"role_id"`
}

type User_Role struct {
	Role_Id int64  `json:"role_id"`
	User_Id int64  `json:"user_id"`
	Role    string `json:"role"`
}
