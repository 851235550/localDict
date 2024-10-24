package main

import (
	"dict/services"
	"flag"
	"fmt"
)

func convertEnglishToChinese(word string) {
	services.NewYoudaoService().E2C(word)
}

func main() {
	eWord := flag.String("e", "", "english word")

	flag.Parse()
	if *eWord != "" {
		fmt.Println(*eWord)
		convertEnglishToChinese(*eWord)
	}
}
