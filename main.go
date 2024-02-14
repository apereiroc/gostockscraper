package main

import (
	"fmt"
	"gostockscrape/argparser"
	"os"
)

func main() {
	parser := argparser.NewArgParser(os.Args[1:])
	fmt.Printf("Parsed args: %s\n", parser.Get())
	fmt.Printf("Is JSON? %t\n", parser.IsJSON())
}
