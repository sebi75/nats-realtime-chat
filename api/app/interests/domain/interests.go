package domain

type Interest struct {
	Id        string `json:"id,omitempty"`
	UserId    string `json:"user_id"`
	Tag       string `json:"tag"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
