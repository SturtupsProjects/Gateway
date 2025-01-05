package entity

type UserUpdateRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Role        string `json:"role"`
	Password    string `json:"password"`
	CompanyId   string `json:"company_id"`
}
type Error struct {
	Message string `json:"message"`
}

type Names struct {
	Name string `json:"name" binding:"required" example:"Electronics" form:"name"`
}

type CreateProductRequest struct {
	CategoryID    string  `json:"category_id" form:"category_id"`
	Name          string  `json:"name" form:"name"`
	BillFormat    string  `json:"bill_format" form:"bill_format"`
	IncomingPrice float64 `json:"incoming_price" form:"incoming_price"`
	StandardPrice float64 `json:"standard_price" form:"standard_price"`
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
	Type     string `json:"type"`
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

type UpdateProductRequest struct {
	CategoryId    string `json:"category_id,omitempty"`
	Name          string `json:"name,omitempty"`
	ImageUrl      string `json:"image_url,omitempty"`
	BillFormat    string `json:"bill_format,omitempty"`
	IncomingPrice int64  `json:"incoming_price,omitempty"`
	StandardPrice int64  `json:"standard_price,omitempty"`
}
type ProductFilter struct {
	CategoryId string `form:"category_id" json:"category_id,omitempty"` // Optional
	Name       string `form:"name" json:"name,omitempty"`               // Optional
	CreatedAt  string `form:"created_at" json:"created_at,omitempty"`   // Optional
	CreatedBy  string `form:"created_by" json:"created_by,omitempty"`   // Optional
	Limit      int64  `form:"limit" json:"limit,omitempty"`             // Optional
	Page       int64  `form:"page" json:"page,omitempty"`               // Optional
}

type FilterPurchase struct {
	ProductId   string `json:"product_id,omitempty"`
	SupplierId  string `json:"supplier_id,omitempty"`
	PurchasedBy string `json:"purchased_by,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
}

type PurchaseUpdate struct {
	Id            string `json:"id,omitempty"`
	SupplierId    string `json:"supplier_id,omitempty"`
	Description   string `json:"description,omitempty"`
	PaymentMethod string `json:"payment_method,omitempty"`
}

type SaleUpdate struct {
	ClientId      string `json:"client_id,omitempty"`
	PaymentMethod string `json:"payment_method,omitempty"`
}
type SaleFilter struct {
	StartDate string `protobuf:"bytes,1,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate   string `protobuf:"bytes,2,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	ClientId  string `protobuf:"bytes,3,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	SoldBy    string `protobuf:"bytes,4,opt,name=sold_by,json=soldBy,proto3" json:"sold_by,omitempty"`
}

type UpdateCategoryRequest struct {
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

type UpdateProductForm struct {
	Name          string `form:"name" binding:"required"`                   // Name of the product
	CategoryId    string `form:"category_id" binding:"required"`            // ID of the product category
	BillFormat    string `form:"bill_format"`                               // Optional billing format
	IncomingPrice int64  `form:"incoming_price" binding:"required,numeric"` // Incoming price of the product
	StandardPrice int64  `form:"standard_price" binding:"required,numeric"` // Standard price of the product
}

type MostSoldProductsRequest struct {
	StartDate string `json:"start_date,omitempty"`
	EndDate   string `json:"end_date,omitempty"`
}
type GetTopEntitiesRequest struct {
	StartDate string `json:"start_date,omitempty"`
	EndDate   string `json:"end_date,omitempty"`
	Limit     int32  `json:"limit,omitempty"`
}

type CreateBulkProductsRequest struct {
	Products []CreateProductRequestBulk
}
type CreateProductRequestBulk struct {
	Name          string  `json:"name,omitempty"`
	BillFormat    string  `json:"bill_format,omitempty"`
	IncomingPrice float64 `json:"incoming_price,omitempty"`
	StandardPrice float64 `json:"standard_price,omitempty"`
	TotalCount    int64   `json:"total_count,omitempty"`
}

type TopClient struct {
	ID       string  `json:"id,omitempty"`
	Name     string  `json:"name"`
	Phone    string  `json:"phone"`
	TotalSum float64 `json:"total_sum"`
}

type TopClientList struct {
	Clients []TopClient `json:"clients"`
}
