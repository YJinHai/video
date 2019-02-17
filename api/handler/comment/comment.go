package comment

type CreateRequest struct {
	VideoID string `json:"video_id" form:"vid"`
	UserID  string `json:"user_id" form:"uid"`
	Content string `json:"content" form:"content"`
}
