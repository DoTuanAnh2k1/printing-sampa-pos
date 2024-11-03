package printing

import "github.com/mect/go-escpos"

func Printing(data string) error {
	p, err := escpos.NewUSBPrinterByPath("")
	if err != nil {
		return err
	}

	err = p.Init()
	if err != nil {
		return err
	}

	p.Print(data)

	p.Feed(2)
	p.Cut()
	p.End()
	return nil
}
