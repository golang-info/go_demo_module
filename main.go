package main

import (
	"github.com/fatih/color"
	"rsc.io/quote"
)

func main() {
	color.Cyan(quote.Hello())
	color.Red(quote.Hello())
}