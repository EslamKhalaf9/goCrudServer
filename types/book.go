package customtypes

type CreateBook struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type GetBook struct {
	ID string `uri:"id" binding:"required"`
}
