package layout

import (
	"printing-sampa-pos/model"
	"strings"
)

func FillValueLayout(ticket model.Ticket, layout string) string {
	return replaceValue(layout, ticket)
}

func replaceValue(layout string, ticket model.Ticket) string {
	// replace mandatory
	layout = strings.ReplaceAll(layout, "{Ticket.Terminal}", ticket.Terminal)
	layout = strings.ReplaceAll(layout, "{Ticket.LoginUser}", ticket.LoginUser)
	layout = strings.ReplaceAll(layout, "{Ticket.PaymentDate}", ticket.PaymentDate)
	layout = strings.ReplaceAll(layout, "{Ticket.PaymentTime}", ticket.PaymentTime)
	layout = strings.ReplaceAll(layout, "{Ticket.PaymentType}", ticket.PaymentType)

	// replace order,
	// includes order, order comp, gift, tag, promotions
	layout = replaceOrder(layout, ticket.Orders)
	layout = replaceOrderComp(layout, ticket.OrderComps)
	layout = replaceOrderTag(layout, ticket.OrderTags)
	layout = replaceOrderPromotion(layout, ticket.OrderPromotion)

	// replace tax
	layout = replaceTax(layout, ticket.Taxes)

	// replace discount
	layout = replaceDiscount(layout, ticket.Discounts)

	// replace service
	layout = replaceService(layout, ticket.Services)

	// replace payment
	layout = replacePayment(layout, ticket.Payments)

	// replace entity
	layout = replaceEntityTable(layout, ticket.Entity.Table)
	layout = replaceEntityMember(layout, ticket.Entity.Member)

	return removeEmptyLines(layout)
}
