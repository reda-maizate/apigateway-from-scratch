package domain

type Note struct {
	Uuid      string `json:"uuid"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedBy string `json:"created_by"`
}
