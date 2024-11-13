package layout

import (
	"printing-sampa-pos/model"
	"regexp"
	"strings"
)

func replaceOrder(layout string, orderList []model.Order) string {
	var orderOut string
	// orderOut = orderOut + "[ORDERS]\n"
	ordersSectionLayout := getOrdersSections(layout)
	for _, order := range orderList {
		quantity := fromStringToInt(order.Quantity)
		price := fromStringToFloat(order.Price)
		totalPrice := float64(quantity) * price
		tmp := ordersSectionLayout
		tmp = strings.ReplaceAll(tmp, "{NAME}", order.Name)
		tmp = strings.ReplaceAll(tmp, "{QUANTITY}", order.Quantity)
		tmp = strings.ReplaceAll(tmp, "{PRICE}", order.Price)
		tmp = strings.ReplaceAll(tmp, "{TOTAL}", fromFloat64ToString(totalPrice, 2))
		tmp += "\n"
		orderOut = orderOut + tmp
	}

	layout = strings.ReplaceAll(layout, "{ORDERS}", orderOut)

	return layout
}

func replaceOrderTag(layout string, orderTagList []model.OrderTag) string {
	var orderTagOut string
	// orderTagOut = orderTagOut + "[ORDERS TAGS]\n"
	ordersTagSectionLayout := getOrderTagSection(layout)
	for _, tag := range orderTagList {
		tmp := ordersTagSectionLayout
		tmp = strings.ReplaceAll(tmp, "{ORDER TAG NAME}", tag.Name)
		tmp = strings.ReplaceAll(tmp, "{ORDER TAG PRICE}", tag.Price)
		tmp += "\n"
		orderTagOut = orderTagOut + tmp
	}
	layout = strings.ReplaceAll(layout, "{ORDERS TAGS}", orderTagOut)

	return layout
}

func replaceOrderComp(layout string, orderCompList []model.OrderComp) string {
	var orderCompOut string
	// orderCompOut = orderCompOut + "[ORDERS:Comp]\n"
	ordersCompSectionLayout := getOrderCompSection(layout)
	for _, comp := range orderCompList {
		tmp := ordersCompSectionLayout
		tmp = strings.ReplaceAll(tmp, "{NAME}", comp.Name)
		tmp = strings.ReplaceAll(tmp, "{QUANTITY}", comp.Quantity)
		tmp = strings.ReplaceAll(tmp, "{PRICE}", comp.Price)
		tmp = strings.ReplaceAll(tmp, "{TOTAL}", comp.Total)
		tmp += "\n"
		orderCompOut = orderCompOut + tmp
	}
	layout = strings.ReplaceAll(layout, "{ORDERS:Comp}", orderCompOut)

	return layout
}

func replaceOrderPromotion(layout string, orderPromotionList []model.OrderPromotion) string {
	var orderPromotionOut string
	// orderPromotionOut = orderPromotionOut + "[ORDERS PROMOTIONS]\n"
	orderPromotionSectionLayout := getOrderPromotionSection(layout)
	for _, comp := range orderPromotionList {
		tmp := orderPromotionSectionLayout
		tmp = strings.ReplaceAll(tmp, "{ORDER PROMOTION NAME}", comp.Name)
		tmp = strings.ReplaceAll(tmp, "{ORDER PROMOTION TOTAL}", comp.Total)
		tmp += "\n"
		orderPromotionOut = orderPromotionOut + tmp
	}

	layout = strings.ReplaceAll(layout, "{ORDER PROMOTIONS}", orderPromotionOut)
	return layout
}

func getOrdersSections(layout string) (ordersSectionLayout string) {
	// get order section
	reOrders := regexp.MustCompile(`(?s)\[ORDERS\](.*?)\[`)
	matchesOrders := reOrders.FindStringSubmatch(layout)
	if len(matchesOrders) > 1 {
		ordersSectionLayout = matchesOrders[1]
	}

	return
}

func getOrderTagSection(layout string) (ordersTagSectionLayout string) {
	// get order tag section
	reOrdersTag := regexp.MustCompile(`(?s)\[ORDER TAGS\](.*?)\[`)
	matchesOrdersTag := reOrdersTag.FindStringSubmatch(layout)
	if len(matchesOrdersTag) > 1 {
		ordersTagSectionLayout = matchesOrdersTag[1]
	}
	return
}

func getOrderCompSection(layout string) (ordersCompSectionLayout string) {
	// get order gifts section
	reOrdersComp := regexp.MustCompile(`(?s)\[ORDERS:Comp\](.*?)\[`)
	matchesOrdersComp := reOrdersComp.FindStringSubmatch(layout)
	if len(matchesOrdersComp) > 1 {
		ordersCompSectionLayout = matchesOrdersComp[1]
	}
	return
}

func getOrderPromotionSection(layout string) (ordersPromotionSectionLayout string) {
	// get order promotion section
	reOrderPromotion := regexp.MustCompile(`(?s)\[ORDER PROMOTIONS\](.*?)\[`)
	matchesOrderPromotion := reOrderPromotion.FindStringSubmatch(layout)
	if len(matchesOrderPromotion) > 1 {
		ordersPromotionSectionLayout = matchesOrderPromotion[1]
	}
	return
}
