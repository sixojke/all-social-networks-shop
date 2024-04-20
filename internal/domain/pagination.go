package domain

type Pagination struct {
	Data       interface{} `json:"data"`
	Limit      int         `json:"limit"`
	TotalItems int         `json:"total_items"`
	TotalPages int         `json:"total_pages"`
}
