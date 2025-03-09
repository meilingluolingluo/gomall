package types

type OrderItem struct {
	Id          int32
	ProductName string
	Picture     string
	Cost        float64
	Quantity    int32
}

type Order struct {
	OrderId       string
	CreateData    string
	OrderItems    []OrderItem
	Cost          float64
	Consignee     Consignee
	PaymentStatus string
}

type Consignee struct {
	Email         string
	StreetAddress string
	City          string
	State         string
	Country       string
	ZipCode       int32
}
