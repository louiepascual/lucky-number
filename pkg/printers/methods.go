package printers

import (
	"fmt"

	"github.com/louiepascual/lucky-number/internal/random"

	"github.com/fatih/color"

)

func PrintLuckyNumber() {
	green := color.New(color.FgGreen)

	fmt.Printf("Your lucky number is ")
	green.Printf("%d!\n", random.Number())

}