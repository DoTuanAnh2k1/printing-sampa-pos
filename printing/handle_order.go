package printing

import (
	"regexp"
	"strings"
)

func HandleOrder(layout string) (string, error) {
	var order strings.Builder

	ordersSectionLayout, ordersGiftSectionLayout, ordersTagSectionLayout := getOrdersSections(layout)
	ordersSectionLayout = strings.ReplaceAll(ordersSectionLayout, "{ORDER TAGS}", ordersTagSectionLayout)
	ordersGiftSectionLayout = strings.ReplaceAll(ordersGiftSectionLayout, "{ORDER TAGS}", ordersTagSectionLayout)
	orderSection, err := handleOrderLine(ordersSectionLayout)
	if err != nil {
		return order.String(), err
	}
	orderGiftSection, err := handleOrderLine(ordersGiftSectionLayout)
	if err != nil {
		return order.String(), err
	}

	order.WriteString(orderSection)
	order.WriteString("\n")
	order.WriteString(orderGiftSection)
	return order.String(), nil
}

func handleOrderLine(orderLine string) (string, error) {
	var orderTag strings.Builder
	lines := strings.Split(orderLine, "\n")
	for _, line := range lines {
		lineFormat, err := GetLineFormat(line)
		if err != nil {
			return orderTag.String(), err
		}
		orderTag.WriteString(lineFormat)
	}
	return removeEmptyLines(orderTag.String()), nil
}

func getOrdersSections(layout string) (ordersSectionLayout string, ordersGiftSectionLayout string, ordersTagSectionLayout string) {
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

	return
}

func removeEmptyLines(input string) string {
	lines := strings.Split(input, "\n")
	var nonEmptyLines []string
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			nonEmptyLines = append(nonEmptyLines, line)
		}
	}
	return strings.Join(nonEmptyLines, "\n")
}
