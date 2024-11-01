package printing

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// Convert from line type string to Line have format
func HandlerLine(line string) (*Line, error) {
	var lineInfo Line

	if strings.Contains(line, "<") && strings.Contains(line, ">") {
		tag, body, height, width, err := HandlerHeader(line)
		if err != nil {
			return nil, err
		}
		if tag == "F" {
			lineInfo.IsFull = true
			lineInfo.Body = HandlerBar(body)
			return &lineInfo, nil
		}
		lineInfo.Body = body
		lineInfo.Format.Align = tag
		if height != -1 {
			lineInfo.Format.FrontHeight = height
		}
		if width != -1 {
			lineInfo.Format.FrontWidth = width
		}
		return &lineInfo, nil
	}

	return &lineInfo, nil
}

func HandlerBar(barType string) string {
	switch barType {
	case "-":
		return Dash
	case "=":
		return DoubleDash
	default:
		ans := ""
		for i := 0; i < (SizeOfReceipt / len(barType)); i++ {
			ans = ans + barType
		}
		return ans
	}
}

// HandlerHeader convert from a line to elements
// It return Tag, Body, Front Height and Front Width
// If Front Height and Front Width does not exsisted, return -1, -1
// Example "<C10>CURRY VILLAGE" -> Center, "CURRY VILLAGE", 1, 0
func HandlerHeader(line string) (string, string, int, int, error) {
	tag := HandlerTag(line)
	parts := strings.Split(line, ">")
	body := parts[len(parts)-1]
	// tag := strings.TrimPrefix(parts[1], "<")
	if len(tag) == 1 {
		return tag, body, -1, -1, nil
	}
	if len(tag) == 3 {
		frontHeight, err := convertFromStringToInt(string(tag[1]))
		if err != nil {
			return "", "", -1, -1, err
		}
		frontWidth, err := convertFromStringToInt(string(tag[2]))
		if err != nil {
			return "", "", -1, -1, err
		}
		return string(tag[0]), body, frontHeight, frontWidth, nil
	}
	return "", "", -1, -1, errors.New("invalid tag")
}

func HandlerTag(line string) string {
	re := regexp.MustCompile(`<([^>]+)>`)

	match := re.FindStringSubmatch(line)
	if len(match) > 1 {
		// fmt.Println(match[1])
		return match[1]
	}
	return ""
}

func convertFromStringToInt(s string) (int, error) {
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0, errors.New("invalid number")
	}
	return num, nil
}