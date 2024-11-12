package phatlc

import (
	layout_utils "printing-sampa-pos/layout"
	"printing-sampa-pos/model"
	"printing-sampa-pos/parsing"
	"printing-sampa-pos/printing"
	"printing-sampa-pos/utils"
)

func PrintFromTicket(ticket model.Ticket, layoutPath string) error {
	layout, err := utils.ReadFromFile(layoutPath)
	if err != nil {
		return err
	}
	layoutFill := layout_utils.FillValueLayout(ticket, layout)

	layoutData, err := printing.HandleLayout(layoutFill)
	if err != nil {
		return err
	}

	err = printing.Printing(layoutData)
	if err != nil {
		return err
	}

	return nil
}

func PrintFromTemplate(templatePath, layoutPath string) error {
	template, err := utils.ReadFromFile(templatePath)
	if err != nil {
		return err
	}

	ticket := parsing.FromTemplateToTicket(template)
	return PrintFromTicket(ticket, layoutPath)
}