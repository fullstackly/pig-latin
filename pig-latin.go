package main

import (
	"flag"
	"fmt"

	"github.com/fullstackly/pig-latin/translator"
)

func main() {
	text := flag.String("translate", "", "translates given string according to pig-latin")

	flag.Parse()

	process(*text)
}

func process(text string) {
	if len(text) != 0 {
		fmt.Println(translator.TranslateText(text))
	} else {
		fmt.Println("Enter command and text. \nEnter -h for help")
	}
}
