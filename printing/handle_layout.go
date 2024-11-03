package printing

import (
	"fmt"
	"strings"
)

func HandleLayout(layout string) (string, error) {
	var output strings.Builder

	// Filter and handle Entities

	// Filter and handle Discounts

	// Filter and handle Orders
	order, err := HandleOrder(layout)
	if err != nil {
		return "", err
	}

	// Combine all into one

	lines := strings.Split(layout, "\n")
	for _, line := range lines {
		lineFormat, err := HandlerLine(line)
		if err != nil {
			return "", err
		}
		fmt.Println(lineFormat)
	}

	output.WriteString(order)

	return output.String(), nil
}
