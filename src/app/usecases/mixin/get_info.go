package mixin

type GetInfoCaseRes struct {
	Id    uint64 `json:"id" example:"1" format:"uint64" copier:"must"`
	Name  string `json:"name" example:"a new dir" format:"string" copier:"must"`
	IsDir bool   `json:"IsDir" example:"true" format:"bool"`
}
