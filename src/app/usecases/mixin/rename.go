package mixin

import "github.com/PowerReport/pfs/src/util/id"

type RenameCaseReq struct {
	Id   id.OpId `json:"id" example:"1" format:"int"`
	Name string  `json:"name" example:"新的目录名称" format:"string"`
}

type RenameCaseRes struct {
	Id    uint64 `json:"id" example:"1" format:"int" copier:"must"`
	Name  string `json:"name" example:"a new dir" format:"string" copier:"must"`
	IsDir bool   `json:"isDir" example:"true" format:"bool"`
}
