package model

type Post struct {
	Id         int32  `json:"id"`
	CategoryId int32  `json:"category"`
	UserId     int32  `json:"user_id"`
	Title      string `json:"title"`
	Body       string `json:"body"`
	Date       string `json:"date"`
}

func (p *Post) IsValid() bool {
	return p.CategoryId != 0 && p.Title != "" && p.Body != ""
}
