package layout

import (
	"fmt"
	"printing-sampa-pos/model"
	"strings"
)

func FillValueLayout(ticket model.Ticket, layout string) string {
	return replaceValue(layout, ticket)
}

func replaceValue(layout string, ticket model.Ticket) string {
	var sb strings.Builder
	// replace mandatory
	layout = replaceTicket(layout, ticket)

	// replace order,
	// includes order, order comp, gift, tag, promotions
	layout = replaceOrder(layout, ticket.Orders)
	layout = replaceOrderComp(layout, ticket.OrderComps)
	layout = replaceOrderTag(layout, ticket.OrderTags)
	layout = replaceOrderPromotion(layout, ticket.OrderPromotion)
	fmt.Println(layout)

	// replace tax
	sb.WriteString(replaceTax(layout, ticket.Taxes))

	// replace discount
	sb.WriteString(replaceDiscount(layout, ticket.Discounts))

	// replace service
	sb.WriteString(replaceService(layout, ticket.Services))

	// replace payment
	sb.WriteString(replacePayment(layout, ticket.Payments))

	// replace entity
	sb.WriteString(replaceEntityTable(layout, ticket.Entity.Table))
	sb.WriteString(replaceEntityMember(layout, ticket.Entity.Member))

	return removeEmptyLines(sb.String())
}
