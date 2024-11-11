package layout

import (
	"printing-sampa-pos/model"
	"regexp"
	"strings"
)

func replaceTax(layout string, taxList []model.Tax) string {
	var taxOut string
	taxOut += "[TAXES]\n"
	taxSectionLayout := getTaxSection(layout)

	for _, tax := range taxList {
		tmp := taxSectionLayout
		tmp = strings.ReplaceAll(tmp, "{TAX NAME}", tax.Name)
		tmp = strings.ReplaceAll(tmp, "{TAX RATE}", tax.Rate)
		tmp = strings.ReplaceAll(tmp, "{TAX AMOUNT}", tax.Amount)
		tmp += "\n"
		taxOut = taxOut + tmp
	}
	return taxOut
}

func getTaxSection(layout string) (taxSectionLayout string) {
	// get payment section
	reTax := regexp.MustCompile(`(?s)\[TAXES\](.*?)\[`)
	matchesTax := reTax.FindStringSubmatch(layout)
	if len(matchesTax) > 1 {
		taxSectionLayout = matchesTax[1]
	}
	return
}
