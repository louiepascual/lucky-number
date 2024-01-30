package main

import (
	"fmt"

	"github.com/louiepascual/lucky-number/internal/random"

	"github.com/fatih/color"

)

func main() {
	green := color.New(color.FgGreen)

	fmt.Printf("Your lucky number is ")
	green.Printf("%d!\n", random.Number())

}

func PrintLuckyNumber() {
	green := color.New(color.FgGreen)

	fmt.Printf("Your lucky number is ")
	green.Printf("%d!\n", random.Number())

}