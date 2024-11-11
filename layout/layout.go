package layout

import (
	"printing-sampa-pos/model"
	"strings"
)

func FillValueLayout(ticket model.Ticket, layout string) string {
	return replaceValue(layout, ticket)
}

func replaceValue(layout string, ticket model.Ticket) string {
	layout = strings.ReplaceAll(layout, "{Ticket.Terminal}", ticket.Terminal)
	layout = strings.ReplaceAll(layout, "{Ticket.LoginUser}", ticket.LoginUser)
	layout = strings.ReplaceAll(layout, "{Ticket.PaymentDate}", ticket.PaymentDate)
	layout = strings.ReplaceAll(layout, "{Ticket.PaymentTime}", ticket.PaymentTime)
	layout = strings.ReplaceAll(layout, "{Ticket.PaymentType}", ticket.PaymentType)

	layout = replaceOrder(layout, ticket.Orders)
	layout = replaceTax(layout, ticket.Taxes)
	layout = replaceDiscount(layout, ticket.Discounts)
	layout = replaceService(layout, ticket.Services)

	return layout
}
