package entity

type UserUpdateRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Role        string `json:"role"`
	Password    string `json:"password"`
}
type Error struct {
	Message string `json:"message"`
}

type Names struct {
	Name string `json:"name" binding:"required" example:"Electronics"`
}

type CreateProductRequest struct {
	CategoryID    string  `json:"category_id"`
	Name          string  `json:"name"`
	BillFormat    string  `json:"bill_format"`
	IncomingPrice float32 `json:"incoming_price"`
	StandardPrice float32 `json:"standard_price"`
}
