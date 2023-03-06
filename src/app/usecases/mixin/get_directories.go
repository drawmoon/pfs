package mixin

type GetDirectoryCaseRes struct {
	Id   uint64 `json:"id" example:"1" format:"uint64" copier:"must"`
	Name string `json:"name" example:"a new dir" format:"string" copier:"must"`
}

type GetDirectoriesCaseRes struct {
	PageIndex    int64                 `json:"pageIndex" example:"1" format:"int"`
	PageSize     int64                 `json:"pageSize" example:"10" format:"int"`
	TotalCount   int64                 `json:"totalCount" example:"50" format:"int"`
	TotalPages   int64                 `json:"totalPages" example:"5" format:"int"`
	Items        []GetDirectoryCaseRes `json:"items" format:"[]GetDirectoryCaseRes"`
	HasPrevPages bool                  `json:"hasPrevPages" example:"true" format:"bool"`
	HasNextPages bool                  `json:"hasNextPages" example:"true" format:"bool"`
}
