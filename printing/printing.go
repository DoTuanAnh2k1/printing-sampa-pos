package printing

import (
	"strings"

	"github.com/mect/go-escpos"
)

func Printing(data string) error {
	lines := strings.Split(data, "\n")
	p, err := escpos.NewUSBPrinterByPath("")
	if err != nil {
		return err
	}

	err = p.Init()
	if err != nil {
		return err
	}

	for _, line := range lines {
		if line == "<EB>" {
			p.EnableBold()
		} else if line == "<DB>" {
			p.DisableBold()
		}
		p.Print(line)
	}

	p.Feed(2)
	p.Cut()
	p.End()
	return nil
}
