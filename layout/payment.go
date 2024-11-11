package layout

import (
	"printing-sampa-pos/model"
	"regexp"
	"strings"
)

func replacePayment(layout string, paymentList []model.Payment) string {
	var paymentOut string
	paymentOut += "[PAYMENTS]\n"
	paymentSectionLayout := getPaymentSection(layout)

	for _, payment := range paymentList {
		tmp := paymentSectionLayout
		tmp = strings.ReplaceAll(tmp, "{PAYMENT NAME}", payment.Name)
		tmp = strings.ReplaceAll(tmp, "{PAYMENT AMOUNT}", payment.Amount)
		tmp += "\n"
		paymentOut = paymentOut + tmp
	}
	return paymentOut
}

func getPaymentSection(layout string) (paymentSectionLayout string) {
	// get payment section
	rePayment := regexp.MustCompile(`(?s)\[PAYMENTS\](.*?)\[`)
	matchesPayment := rePayment.FindStringSubmatch(layout)
	if len(matchesPayment) > 1 {
		paymentSectionLayout = matchesPayment[1]
	}
	return
}
