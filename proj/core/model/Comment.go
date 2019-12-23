package model

type Comment struct {
	Id     int32
	PostId int32
	UserId int32
	Body   string
}
