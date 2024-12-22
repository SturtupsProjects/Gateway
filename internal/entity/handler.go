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
	CategoryID    string `json:"category_id"`
	Name          string `json:"name"`
	BillFormat    string `json:"bill_format"`
	IncomingPrice int64  `json:"incoming_price"`
	StandardPrice int64  `json:"standard_price"`
}

type Sale struct {
	ClientId      string       `json:"client_id,omitempty"`
	PaymentMethod string       `json:"payment_method,omitempty"`
	SoldProducts  []*SalesItem `json:"sold_products,omitempty"`
}

type SalesItem struct {
	Id         string  `json:"id,omitempty"`
	SaleId     string  `json:"sale_id,omitempty"`
	ProductId  string  `json:"product_id,omitempty"`
	Quantity   int32   `json:"quantity,omitempty"`
	SalePrice  float64 `json:"sale_price,omitempty"`
	TotalPrice float64 `json:"total_price,omitempty"`
}

type Purchase struct {
	SupplierId    string          `json:"supplier_id,omitempty"`
	Description   string          `json:"description,omitempty"`
	PaymentMethod string          `json:"payment_method,omitempty"`
	Items         []*PurchaseItem `json:"items,omitempty"`
}

type PurchaseItem struct {
	ProductId     string  `json:"product_id,omitempty"`
	Quantity      int32   `json:"quantity,omitempty"`
	PurchasePrice float64 `json:"purchase_price,omitempty"`
}

type Client struct {
	FullName string `json:"full_name"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
}

type CreateCompanyRequest struct {
	Name    string `json:"name"`
	Website string `json:"website"`
	Logo    string `json:"logo"`
}
type CreateUserToCompanyRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	Role        string `json:"role" default:"user"`
}
type UpdateCompanyRequest struct {
	Name    string `json:"name"`
	Website string `json:"website"`
	Logo    string `json:"logo"`
}
