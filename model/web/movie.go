package web

type MovieCreateRequestBody struct {
	Title string `json:"title" binding:"required"`
	Year  string `json:"year" binding:"required"`
}

type MovieCreateResponseBody struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
