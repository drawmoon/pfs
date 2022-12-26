package mixin

type GetFileCaseRes struct {
}

type GetFilesCaseRes struct {
	PageIndex    int              `json:"pageIndex" example:"1" format:"int"`
	PageSize     int              `json:"pageSize" example:"10" format:"int"`
	TotalCount   int              `json:"totalCount" example:"50" format:"int"`
	TotalPages   int              `json:"totalPages" example:"5" format:"int"`
	Items        []GetFileCaseRes `json:"items" format:"[]GetFileCaseRes"`
	HasPrevPages bool             `json:"hasPrevPages" example:"true" format:"bool"`
	HasNextPages bool             `json:"hasNextPages" example:"true" format:"bool"`
}
