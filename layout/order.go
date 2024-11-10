package layout

import (
	"printing-sampa-pos/model"
	"regexp"
	"strings"
)

func replaceOrder(layout string, orderList []model.Order) string {
	var orderOut string
	orderOut = orderOut + "[ORDERS]\n"
	ordersSectionLayout, _, _, _ := getOrdersSections(layout)
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
	return orderOut
}

func getOrdersSections(layout string) (
	ordersSectionLayout string,
	ordersGiftSectionLayout string,
	ordersTagSectionLayout string,
	ordersPromotionSectionLayout string,
) {
	// get order section
	reOrders := regexp.MustCompile(`(?s)\[ORDERS\](.*?)\[`)
	matchesOrders := reOrders.FindStringSubmatch(layout)
	if len(matchesOrders) > 1 {
		ordersSectionLayout = matchesOrders[1]
	}

	// get order gifts section
	reOrdersGifts := regexp.MustCompile(`(?s)\[ORDERS:Gift\](.*?)\[`)
	matchesOrdersGifts := reOrdersGifts.FindStringSubmatch(layout)
	if len(matchesOrdersGifts) > 1 {
		ordersGiftSectionLayout = matchesOrdersGifts[1]
	}

	// get order tag section
	reOrdersTag := regexp.MustCompile(`(?s)\[ORDER TAGS\](.*?)\[`)
	matchesOrdersTag := reOrdersTag.FindStringSubmatch(layout)
	if len(matchesOrdersTag) > 1 {
		ordersTagSectionLayout = matchesOrdersTag[1]
	}

	// get order promotion section
	reOrderPromotion := regexp.MustCompile(`(?s)\[ORDER PROMOTIONS\](.*?)\[`)
	matchesOrderPromotion := reOrderPromotion.FindStringSubmatch(layout)
	if len(matchesOrderPromotion) > 1 {
		ordersPromotionSectionLayout = matchesOrderPromotion[1]
	}

	return
}
