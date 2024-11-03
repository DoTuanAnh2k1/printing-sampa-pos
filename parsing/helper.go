package parsing

import (
	"printing-sampa-pos/utils"
	"strings"
)

func getTemplate(inputFileTemplatePath string) string {
	lines, err := utils.ReadFromFileToLines(inputFileTemplatePath)
	utils.CheckErr(err)

	m := make(map[string]string)
	for _, line := range lines {
		if strings.Contains(line, ":") {
			parts := strings.Split(line, ":")
			m[parts[0]] = parts[1]

		}
	}
	return lines[0]
}

func getLayout(inputFileLayoutPath string) string {
	layout, err := utils.ReadFromFile(inputFileLayoutPath)
	utils.CheckErr(err)
	return layout
}
