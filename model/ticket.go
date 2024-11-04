package model

type Ticket struct {
	Terminal    string
	Cashier     string // Doesn't require
	PaymentDate string
	PaymentType string
	Tag         Tag
	Payments    []Payment
	Orders      []Order
}

type Tag struct {
	Pax     int
	PaxTime string
}

type Payment struct {
	Name               string
	Tendered           string
	PaymentInformation PaymentInfo
}

type PaymentInfo struct {
	RefNo   int
	RefTime string
}

type Order struct {
	Name     string
	Quantity string
	Price    string
}
