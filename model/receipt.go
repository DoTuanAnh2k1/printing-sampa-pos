package model

import (
	"fmt"
	"strings"
	"time"
)

type ReceiptLayOut struct {
	Header      HeaderReceipt
	ReceiptInfo ReceiptInfo
	Entities    Entities
}

type HeaderReceipt struct {
	Something   string
	StoreName   string
	Address     string
	PhoneNumber string
}

type ReceiptInfo struct {
	ReceiptNo int
	Date      time.Time
}

type Entities struct {
	EntityName  string    // Table Name
	Header      [4]string // asumming only 4
	Orders      []OrderInfo
	PaymentInfo PaymentInfor
}

type OrderInfo struct {
	Quantity    int
	Name        string
	Price       float64
	TotalAmount float64
	OrderTags   []OrderTag
}

type OrderGift struct {
	Quantity  int
	Name      string
	Gift      string
	OrderTags []OrderTag
}

type OrderTag struct {
	Name  string
	Price string
}

type Discount struct {
	Name   string
	Amount string
	Total  string
}

// PaymentInfor for Receipt
type PaymentInfor struct {
	Name   string
	Amount float64
}

func (r ReceiptLayOut) Output() string {
	var sb strings.Builder
	sb.WriteString(CenterlizeReceipt(strings.ToUpper(r.Header.Something)))
	sb.WriteString(CenterlizeReceipt(strings.ToUpper(r.Header.StoreName)))
	sb.WriteString(CenterlizeReceipt(strings.ToUpper(r.Header.Address)))
	sb.WriteString(CenterlizeReceipt("TEL : " + r.Header.PhoneNumber))

	sb.WriteString(Dash)
	sb.WriteString("\n")

	sb.WriteString(CenterlizeReceipt(fmt.Sprintf("Receipt No: %d", r.ReceiptInfo.ReceiptNo)))
	sb.WriteString(JustifyRecept("Date:", r.ReceiptInfo.Date.Format("2006-01-02 15:04:05")))

	sb.WriteString(r.Entities.Output())

	sb.WriteString(JustifyRecept("Total:", fmt.Sprintf("%0.2f", r.Entities.PaymentInfo.Amount)))
	return sb.String()
}

func (e Entities) Output() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("\nTable: %s\n", e.EntityName))
	sb.WriteString(Dash)
	sb.WriteString("\n")
	sb.WriteString(fmt.Sprintf("%-5s%-27s%-10s%-10s\n",
		e.Header[0], e.Header[1], e.Header[2], e.Header[3]))
	for _, order := range e.Orders {
		sb.WriteString(fmt.Sprintf("%-5d%-27s%-10.2f%-10.2f\n",
			order.Quantity, order.Name, order.Price, order.TotalAmount))
	}
	sb.WriteString(DoubleDash)
	sb.WriteString("\n")
	return sb.String()
}

func NewReceipt() ReceiptLayOut {
	return ReceiptLayOut{
		Header: HeaderReceipt{
			Something:   "Curry Village",
			StoreName:   "Banana leaf p/l",
			Address:     "8 Lim Teck Kim Road",
			PhoneNumber: "6226 2562",
		},
		ReceiptInfo: ReceiptInfo{
			ReceiptNo: 3,
			Date:      time.Now(),
		},
		Entities: Entities{
			EntityName: "B14",
			Header:     [4]string{"Qty", "Items", "Price", "Amount"},
			Orders: []OrderInfo{
				{
					Quantity:    1,
					Name:        "Toasted Bagel Jam",
					Price:       1.50,
					TotalAmount: 1.50,
				},
				{
					Quantity:    1,
					Name:        "Toast and Jam",
					Price:       1.50,
					TotalAmount: 1.50,
				},
				{
					Quantity:    1,
					Name:        "Bacon and Tomato",
					Price:       3.49,
					TotalAmount: 3.49,
				},
				{
					Quantity:    100,
					Name:        "Egg, Bacon Cheese",
					Price:       3.99,
					TotalAmount: 399.00,
				},
			},
			PaymentInfo: PaymentInfor{
				Name:   "Total",
				Amount: 405.49,
			},
		},
	}
}
