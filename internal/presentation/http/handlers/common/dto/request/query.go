package request

type PaginationQuery struct {
	Page  int    `form:"page,default=1" binding:"omitempty,min=1"`
	Limit int    `form:"limit,default=10" binding:"omitempty,min=1,max=100"`
	Sort  string `form:"sort"`
}
