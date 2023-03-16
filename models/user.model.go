package models

type User struct {
	CID      string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func (c User) ID() (jsonField string, value interface{}) {
	value = c.CID
	jsonField = "id"
	return
}
