package layout

import (
	"printing-sampa-pos/model"
	"strings"
)

func replaceTicket(layout string, ticket model.Ticket) string {
	layout = strings.ReplaceAll(layout, "{TICKET NO}", ticket.Index)
	layout = strings.ReplaceAll(layout, "{CURRENT TERMINAL}", ticket.Terminal)
	layout = strings.ReplaceAll(layout, "{Ticket.LoginUser}", ticket.LoginUser)
	layout = strings.ReplaceAll(layout, "{TICKET PAYMENT DATE}", ticket.PaymentDate)
	layout = strings.ReplaceAll(layout, "{TICKET PAYMENT TIME}", ticket.PaymentTime)
	layout = strings.ReplaceAll(layout, "{Ticket.PaymentType}", ticket.PaymentType)

	return layout
}
