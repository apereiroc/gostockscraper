package argparser

import (
	"fmt"
	"os"
	"strings"
)

type ArgParser struct {
	args   []string
	isJson bool
}

func NewArgParser(args []string) *ArgParser {
	parser := &ArgParser{args: args, isJson: false}
	parser.parse()
	return parser
}

func (parser ArgParser) parse() {
	// Check for the correct number of arguments
	if len(parser.args) != 1 {
		fmt.Println("Usage: ./main <arg>")
		fmt.Println("<arg> is company symbol (e.g., AMZN) or a path to a json file.")
		os.Exit(0)
	}

	if strings.HasSuffix(parser.args[0], ".json") {
		parser.isJson = true
	}
}
