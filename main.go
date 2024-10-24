package main

import (
	"dict/services"
	"flag"
	"fmt"
)

func convertEnglishToChinese(word string) {
	if len(word) == 0 {
		return
	}
	services.NewYoudaoService().E2C(word)
}

func main() {
	var eWord string
	flag.StringVar(&eWord, "e", "", "Input an English word")

	flag.Parse()
	if eWord != "" {
		fmt.Println(eWord)
		convertEnglishToChinese(eWord)
	}
}
