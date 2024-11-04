package main

import (
	"fmt"
	"printing-sampa-pos/printing"
)

func main() {
	output, err := printing.HandleLayout(printing.LayoutTest)
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println(output)
}
