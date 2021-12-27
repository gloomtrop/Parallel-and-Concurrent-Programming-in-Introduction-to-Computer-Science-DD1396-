package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	wordcount := make(map[string]int)
	
	for i := 0; i<len(words); i++{
		_,ok := wordcount[words[i]]
		if ok == false{
			wordcount[words[i]] = 1
		} else{
			wordcount[words[i]] +=1
		}
	}
	return wordcount
}

func main() {
	wc.Test(WordCount)
}