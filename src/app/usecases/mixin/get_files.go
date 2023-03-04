package mixin

type GetFileCaseRes struct {
}

type GetFilesCaseRes struct {
	PageIndex    int64            `json:"pageIndex" example:"1" format:"int"`
	PageSize     int64            `json:"pageSize" example:"10" format:"int"`
	TotalCount   int64            `json:"totalCount" example:"50" format:"int"`
	TotalPages   int64            `json:"totalPages" example:"5" format:"int"`
	Items        []GetFileCaseRes `json:"items" format:"[]GetFileCaseRes"`
	HasPrevPages bool             `json:"hasPrevPages" example:"true" format:"bool"`
	HasNextPages bool             `json:"hasNextPages" example:"true" format:"bool"`
}
