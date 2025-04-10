package entity

import "gateway/internal/generated/products"

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
	Quantity      int64   `json:"quantity" form:"quantity"`
}

type Sale struct {
	ClientId      string       `json:"client_id,omitempty"`
	PaymentMethod string       `json:"payment_method,omitempty"`
	ClientName    string       `json:"client_name,omitempty"`
	ClientPhone   string       `json:"client_phone,omitempty"`
	IsForDebt     bool         `json:"is_for_debt,omitempty" default:"false"`
	PaidAmount    float64      `json:"paid_amount,omitempty"`
	SoldProducts  []*SalesItem `json:"sold_products,omitempty"`
}
type PaymentSale struct {
	ClientId      string                `json:"client_id,omitempty"`
	PaymentMethod string                `json:"payment_method,omitempty"`
	IsFullyDebt   bool                  `json:"is_fully_debt,omitempty"`
	CurrencyCode  string                `json:"currency_code,omitempty"`
	PaidAmount    float64               `json:"paid_amount,omitempty"`
	BranchId      string                `json:"branch_id,omitempty"`
	SoldProducts  []*products.SalesItem `json:"sold_products,omitempty"`
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
	FullName   string `json:"full_name"`
	Address    string `json:"address"`
	Phone      string `json:"phone"`
	ClientType string `json:"client_type"`
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
	TotalCount int64  `form:"total_count" json:"total_count,omitempty"`
}

type FilterPurchase struct {
	ProductName string  `protobuf:"bytes,1,opt,name=product_name,json=product_name,proto3" json:"product_name,omitempty"`
	SupplierId  string  `protobuf:"bytes,2,opt,name=supplier_id,json=supplierId,proto3" json:"supplier_id,omitempty"`
	PurchasedBy string  `protobuf:"bytes,3,opt,name=purchased_by,json=purchasedBy,proto3" json:"purchased_by,omitempty"`
	CreatedAt   string  `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Limit       int64   `protobuf:"varint,7,opt,name=limit,proto3" json:"limit,omitempty"`
	Page        int64   `protobuf:"varint,8,opt,name=page,proto3" json:"page,omitempty"`
	Description string  `protobuf:"bytes,9,opt,name=description,proto3" json:"description,omitempty"`
	TotalCost   float64 `json:"total_cost,omitempty"`
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
	StartDate string `json:"start_date,omitempty"` // Дата начала для фильтрации
	EndDate   string `json:"end_date,omitempty"`   // Дата окончания для фильтрации
	ClientId  string `json:"client_id,omitempty"`  // ID клиента для фильтрации
	SoldBy    string `json:"sold_by,omitempty"`    // ID продавца для фильтрации
	CompanyId string `json:"company_id,omitempty"` // ID компании для фильтрации
	Limit     int64  `json:"limit,omitempty"`      // Количество записей для возврата (по умолчанию 10)
	Page      int64  `json:"page,omitempty"`       // Номер страницы для пагинации (по умолчанию 1)
	BranchId  string `json:"branch_id,omitempty"`
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

type StreetClientFilter struct {
	Limit int32  `json:"limit,omitempty"`
	Page  int32  `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	Name  string `json:"name,omitempty"`
	Phone string `json:"phone,omitempty"`
}

type ClientFilter struct {
	FullName   string `protobuf:"bytes,1,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Address    string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Phone      string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Page       int32  `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	Limit      int32  `protobuf:"varint,5,opt,name=limit,proto3" json:"limit,omitempty"`
	ClientType string `json:"client_type,omitempty"`
}

type TransferReq struct {
	ToBranchId  string                  `protobuf:"bytes,3,opt,name=to_branch_id,json=toBranchId,proto3" json:"to_branch_id,omitempty"`
	Description string                  `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Products    []*TransfersProductsReq `protobuf:"bytes,5,rep,name=products,proto3" json:"products,omitempty"`
	CompanyId   string                  `protobuf:"bytes,6,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
}

type TransfersProductsReq struct {
	ProductId       string `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	ProductQuantity int64  `protobuf:"varint,2,opt,name=product_quantity,json=productQuantity,proto3" json:"product_quantity,omitempty"`
}

type SalaryRequest struct {
	UserId       string  `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CurrencyCode string  `protobuf:"bytes,2,opt,name=currency_code,json=currencyCode,proto3" json:"currency_code,omitempty"`
	Amount       float64 `protobuf:"fixed64,3,opt,name=amount,proto3" json:"amount,omitempty"`
	SalaryDate   string  `protobuf:"bytes,4,opt,name=salary_date,json=salaryDate,proto3" json:"salary_date,omitempty"`
}

type SalaryUpdate struct {
	CurrencyCode string  `protobuf:"bytes,2,opt,name=currency_code,json=currencyCode,proto3" json:"currency_code,omitempty"`
	Amount       float64 `protobuf:"fixed64,3,opt,name=amount,proto3" json:"amount,omitempty"`
	SalaryDate   string  `protobuf:"bytes,4,opt,name=salary_date,json=salaryDate,proto3" json:"salary_date,omitempty"`
}

type AdjustmentRequest struct {
	UserId         string  `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	AdjustmentType string  `protobuf:"bytes,2,opt,name=adjustment_type,json=adjustmentType,proto3" json:"adjustment_type,omitempty"`
	CurrencyCode   string  `protobuf:"bytes,3,opt,name=currency_code,json=currencyCode,proto3" json:"currency_code,omitempty"`
	Amount         float64 `protobuf:"fixed64,4,opt,name=amount,proto3" json:"amount,omitempty"`
	AdjustmentDate string  `protobuf:"bytes,5,opt,name=adjustment_date,json=adjustmentDate,proto3" json:"adjustment_date,omitempty"`
}

type AdjustmentUpdate struct {
	AdjustmentType string  `protobuf:"bytes,2,opt,name=adjustment_type,json=adjustmentType,proto3" json:"adjustment_type,omitempty"`
	CurrencyCode   string  `protobuf:"bytes,3,opt,name=currency_code,json=currencyCode,proto3" json:"currency_code,omitempty"`
	Amount         float64 `protobuf:"fixed64,4,opt,name=amount,proto3" json:"amount,omitempty"`
	AdjustmentDate string  `protobuf:"bytes,5,opt,name=adjustment_date,json=adjustmentDate,proto3" json:"adjustment_date,omitempty"`
}

type ClientRequest struct {
	FullName string `json:"full_name,omitempty"`
	Address  string `json:"address,omitempty"`
	Phone    string `json:"phone,omitempty"`
}

type SupplierFilter struct {
	FullName string `protobuf:"bytes,1,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Address  string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Phone    string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Page     int32  `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	Limit    int32  `protobuf:"varint,5,opt,name=limit,proto3" json:"limit,omitempty"`
}

type DebtsRequest struct {
	ClientId     string  `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	TotalAmount  float64 `protobuf:"fixed64,3,opt,name=total_amount,json=totalAmount,proto3" json:"total_amount,omitempty"`
	CurrencyCode string  `protobuf:"bytes,4,opt,name=currency_code,json=currencyCode,proto3" json:"currency_code,omitempty"`
	ShouldPayAt  string  `protobuf:"bytes,6,opt,name=should_pay_at,json=shouldPayAt,proto3" json:"should_pay_at,omitempty"`
}

type PayDebtReq struct {
	DebtId     string  `protobuf:"bytes,1,opt,name=debt_id,json=debtId,proto3" json:"debt_id,omitempty"`
	PaidAmount float64 `protobuf:"fixed64,3,opt,name=paid_amount,json=paidAmount,proto3" json:"paid_amount,omitempty"`
}
