package main

import (
	"flag"
	"fmt"

	"github.com/fullstackly/pig-latin/translator"
)

var (
	text = flag.String("translate", "", "translates given string according to pig-latin")
)

func main() {
	flag.Parse()
	process()
}

func process() {
	if len(*text) != 0 {
		fmt.Println(translator.TranslateText(*text))
	} else {
		fmt.Println("Enter command and text. \nEnter -h for help")
	}
}
