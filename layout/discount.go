package layout

import (
	"printing-sampa-pos/model"
	"regexp"
	"strings"
)

func replaceDiscount(layout string, discountList []model.Discount) string {
	var discountOut string
	discountOut += "[DISCOUNTS]\n"
	discountLayout := getDiscountSection(layout)

	for _, discount := range discountList {
		tmp := discountLayout
		tmp = strings.ReplaceAll(tmp, "{CALCULATION NAME}", discount.Name)
		tmp = strings.ReplaceAll(tmp, "{CALCULATION DESCRIPTION}", discount.Description)
		tmp = strings.ReplaceAll(tmp, "{CALCULATION TOTAL}", discount.Total)
		tmp += "\n"
		discountOut = discountOut + tmp
	}

	return discountOut
}

func getDiscountSection(layout string) (discountSectionLayout string) {
	// get order section
	reDiscounts := regexp.MustCompile(`(?s)\[DISCOUNTS\](.*?)\[`)
	matchesDiscounts := reDiscounts.FindStringSubmatch(layout)
	if len(matchesDiscounts) > 1 {
		discountSectionLayout = matchesDiscounts[1]
	}

	return
}
