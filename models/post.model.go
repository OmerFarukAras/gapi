package models

type Post struct {
	CID       string `json:"id"`
	Author    string `json:"authorid"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdat"`
	UpdatedAt string `json:"updatedat"`
}

func (c Post) ID() (jsonField string, value interface{}) {
	value = c.CID
	jsonField = "id"
	return
}
