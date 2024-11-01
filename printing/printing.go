package printing

import "github.com/mect/go-escpos"

func NewPrinter() (*escpos.Printer, error) {
	p, err := escpos.NewUSBPrinterByPath("")
	if err != nil {
		return nil, err
	}
	return p, nil
}

func ParsingForPrinting(p *escpos.Printer, lines []string) error {
	err := p.Init()
	if err != nil {
		return err
	}

	// for _, line := range lines {

	// }

	p.Feed(2)
	p.Cut()
	p.End()
	return nil
}
