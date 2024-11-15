package pos

import (
	layout_utils "github.com/DoTuanAnh2k1/printing-sampa-pos/layout"
	"github.com/DoTuanAnh2k1/printing-sampa-pos/model"
	"github.com/DoTuanAnh2k1/printing-sampa-pos/parsing"
	"github.com/DoTuanAnh2k1/printing-sampa-pos/printing"
	"github.com/DoTuanAnh2k1/printing-sampa-pos/utils"
)

func PrintFromTicket(ticket model.Ticket, layoutPath string) error {
	layout, err := utils.ReadFromFile(layoutPath)
	if err != nil {
		return err
	}
	layoutFill := layout_utils.FillValueLayout(ticket, layout)

	layoutData, commandList, err := printing.HandleLayout(layoutFill)
	if err != nil {
		return err
	}

	err = printing.Printing(layoutData, commandList)
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
