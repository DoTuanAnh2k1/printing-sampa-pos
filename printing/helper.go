package printing

import (
	"errors"
	"strconv"
)

func convertFromStringToInt(s string) (int, error) {
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0, errors.New("invalid number")
	}
	return num, nil
}
