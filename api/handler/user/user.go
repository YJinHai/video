package user

type CreateRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"pwd"`
}
