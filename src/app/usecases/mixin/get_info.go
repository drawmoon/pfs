package mixin

type GetInfoCaseRes struct {
	Id          uint64 `json:"id" example:"1" format:"uint64"`
	Name        string `json:"name" example:"a new dir" format:"string"`
	IsDirectory bool   `json:"isDirectory" example:"true" format:"bool"`
}
