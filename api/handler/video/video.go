package video

type CreateRequest struct {
	Title        string         `json:"title" form:"title"`
	UserID       string         `json:"user_id" form:"uid"`
	//DisplayCtime sql.NullString `json:"display_ctime" form:"dc"`
}
