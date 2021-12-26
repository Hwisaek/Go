package main

import (
	"fmt"

	"github.com/Hwisaek/Go/nomadcoder_go-for-beginners/mydict"
)

func main() {
	dictionary := mydict.Dictioinary{"first": "First word"}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(word)
}
