package dto

type PageRequest struct {
	Page int
	Size int
}

type Paging struct {
	Page       int `json:"page"`
	Size       int `json:"size"`
	TotalRows  int `json:"totalRows"`
	TotalPages int `json:"totalPages"`
}
