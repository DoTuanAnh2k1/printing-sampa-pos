package main

import (
	"fmt"
	layout_utils "printing-sampa-pos/layout"
	"printing-sampa-pos/parsing"
	"printing-sampa-pos/printing"
	"printing-sampa-pos/utils"
)

func main() {
	templatePath := "data/template/template.txt"
	layoutPath := "data/layout/layout.txt"

	template, err := utils.ReadFromFile(templatePath)
	if err != nil {
		panic(err)
	}

	ticket := parsing.FromTemplateToTicket(template)

	layout, err := utils.ReadFromFile(layoutPath)
	if err != nil {
		panic(err)
	}

	layoutFill := layout_utils.FillValueLayout(ticket, layout)
	// fmt.Println(layoutFill)

	layoutData, err := printing.HandleLayout(layoutFill)
	if err != nil {
		panic(err)
	}

	fmt.Println(layoutData)
}
