package models

type GetListRequest struct {
	Page    int64  `json:"page" binding:"required"`
	PerPage int64  `json:"perPage" binding:"required"`
	Keyword string `json:"keyword"`
}

type SubmitTodosRequest struct {
	Id        string `json:"id"`
	Title     string `json:"title" binding:"required"`
	Completed bool   `json:"complete"`
}

type GetListResponse struct {
	Data   []Todo  `json:"data"`
}