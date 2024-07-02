package dto

type UserRequest struct {
	UserID   int    `uri:"user_id"`
	Username string `form:"username"`
	Page     int    `form:"page"`
	Size     int    `form:"size"`
}
