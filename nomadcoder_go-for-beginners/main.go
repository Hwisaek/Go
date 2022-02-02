package main

import (
	"fmt"

	"github.com/Hwisaek/Go/nomadcoder_go-for-beginners/mydict"
)

func main() {
	dictionary := mydict.Dictioinary{"first": "First word"}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	dictionary.Update(baseWord, "Second")
	word, _ := dictionary.Search(baseWord)
	fmt.Println(word)
}
