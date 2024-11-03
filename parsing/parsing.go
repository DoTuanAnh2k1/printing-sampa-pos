package parsing

func ParsingValue(inputFileLayoutPath, inputFileTemplatePath string) string {
	template := getTemplate(inputFileTemplatePath)
	layout := getLayout(inputFileLayoutPath)

	// return first
	return template + layout
}
