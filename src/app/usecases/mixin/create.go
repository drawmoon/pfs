package mixin

type CreateCaseReq struct {
	Name        string `json:"name" example:"新建目录" format:"string"`
	DirectoryId uint64 `json:"directoryId" example:"1" format:"int"`
	Type        string `json:"type" example:"directory" format:"string"`
}

type CreateCaseRes struct {
	Name        string `json:"name" example:"新建目录" format:"string"`
	DirectoryId uint64 `json:"directoryId" example:"1" format:"int"`
	IsDir       bool   `json:"IsDir" example:"true" format:"bool"`
}
