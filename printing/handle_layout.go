package printing

import (
	"regexp"
	"strings"
)

func HandleLayout(layoutFull string) (string, error) {
	var output strings.Builder

	// Get Main Layout
	layoutSection := getLayout(layoutFull)
	layout, err := formatLayout(layoutSection)
	if err != nil {
		return "", err
	}

	// Filter and handle Entities
	entityFormat, err := HandleEntity(layoutFull)
	if err != nil {
		return "", err
	}

	// Filter and handle Discounts
	discountFormat, err := HandleDiscount(layoutFull)
	if err != nil {
		return "", err
	}

	// Filter and handle Orders
	orderFormat, err := HandleOrder(layoutFull)
	if err != nil {
		return "", err
	}

	// Filter and handle payment
	paymentFormat, err := HandlePayment(layoutFull)
	if err != nil {
		return "", err
	}

	// Filter and handle service
	serviceFormat, err := HandleService(layoutFull)
	if err != nil {
		return "", err
	}

	// Filter and handle tax
	taxFormat, err := HandleTax(layoutFull)
	if err != nil {
		return "", err
	}

	// Combine all into one
	layout = strings.ReplaceAll(layout, "{ENTITIES}", entityFormat)
	layout = strings.ReplaceAll(layout, "{DISCOUNTS}", discountFormat)
	layout = strings.ReplaceAll(layout, "{ORDERS}", orderFormat)
	layout = strings.ReplaceAll(layout, "{PAYMENTS}", paymentFormat)
	layout = strings.ReplaceAll(layout, "{SERVICES}", serviceFormat)
	layout = strings.ReplaceAll(layout, "{TAXES}", taxFormat)

	output.WriteString(layout)

	return output.String(), nil
}

func getLayout(layoutFull string) (layout string) {
	// get entity table section
	reLayOut := regexp.MustCompile(`(?s)\[LAYOUT\](.*?)\[`)
	matchesLayout := reLayOut.FindStringSubmatch(layoutFull)
	if len(matchesLayout) > 1 {
		layout = matchesLayout[1]
	}
	return
}

func formatLayout(layout string) (string, error) {
	var layoutOutput strings.Builder
	lines := strings.Split(layout, "\n")
	for _, line := range lines {
		lineFormat, err := GetLineFormat(line)
		if err != nil {
			return layoutOutput.String(), err
		}
		layoutOutput.WriteString(lineFormat)
	}
	return removeEmptyLines(layoutOutput.String()), nil
}
