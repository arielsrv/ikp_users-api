package model

type CreateUserRequest struct {
	Nickname string `json:"nickname,omitempty"`
	Email    string `json:"email,omitempty"`
}

type CreateUserResponse struct {
	ID int64 `json:"id,omitempty"`
}
