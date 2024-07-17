package web

type MovieCreateRequestBody struct {
	Title string `json:"title" binding:"required"`
	Year  string `json:"year" binding:"required"`
}

type MovieUpdateRequestBody struct {
	Title string `json:"title" binding:"required"`
	Year  string `json:"year" binding:"required"`
}
