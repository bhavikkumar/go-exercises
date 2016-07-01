package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	wordMap := make(map[string]int)
	words := strings.Fields(s)
	for _, i := range words {
		wordMap[i]++
	}
	
	return wordMap
}

func main() {
	wc.Test(WordCount)
}
