package mixin

type CreateCaseReq struct {
	Name        string `json:"name" example:"新建目录" format:"string"`
	DirectoryId uint64 `json:"directoryId" example:"1" format:"int"`
	Type        string `json:"type" example:"directory" format:"string"`
}

type CreateCaseRes struct {
	Id          uint64 `json:"id" example:"1" format:"uint64" copier:"must"`
	Name        string `json:"name" example:"新建目录" format:"string" copier:"must"`
	DirectoryId uint64 `json:"directoryId" example:"1" format:"int" copier:"must"`
	IsDir       bool   `json:"IsDir" example:"true" format:"bool"`
}
