package parsing

var TemplateTest = `
Terminal        : Terminal
Cashier         :  Chronical Do
Date            :   2024-11-03 16:59:34
Bill            : 2024-11-03 16:59:40
[Cover:10]
#Ticket.Orders#
#Ticket.Discounts#
#Ticket.Services#
#Ticket.Taxes#
#Ticket.Payments#
##Ticket.Payments##
Tendered    :   gyxllurslu
Change      :   ixzsyhwrwe
RefNo       :   97
Tendered    :   ytmertfwjc
Change      :   pkddorsvgl
RefNo       :   717
##Ticket.Orders##
Name Highland Coffee [1.00] [15.00]
Name Red Bull [1.00] [18.50]
`
