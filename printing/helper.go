package printing

import (
	"errors"
	"strconv"
	"strings"
)

func convertFromStringToInt(s string) (int, error) {
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0, errors.New("invalid number")
	}
	return num, nil
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
