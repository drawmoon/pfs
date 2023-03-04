package mixin

import "github.com/PowerReport/pfs/src/util/identity"

type DeleteCaseReq struct {
	Id identity.OpId `json:"id" example:"1" format:"int"`
}

type DeleteCaseRes struct {
	Id    uint64 `json:"id" example:"1" format:"int" copier:"must"`
	Name  string `json:"name" example:"a new dir" format:"string" copier:"must"`
	IsDir bool   `json:"isDir" example:"true" format:"bool"`
}
