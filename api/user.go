package api

type createUserRequest struct {
	FirstName string `json:"first_name" binding:"required,min=2,max=16"`
	LastName  string `json:"last_name" binding:"required,min=2,max=16"`
	Username  string `json:"username" binding:"required,min=4,max=16"`
	Nic       string `json:"nic" binding:"required,min=10,max=16"`
	Password  string `json:"password" binding:"required,min=8,max=16"`
	Email     string `json:"email" binding:"required,email,min=16,max=48,endswith=@gmail.com @outlook.com"`
	Phone     string `json:"phone" binding:"required,min=11,max=13"`
}
