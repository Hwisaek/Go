package main

import (
	"fmt"

	"github.com/Hwisaek/Go/nomadcoder_go-for-beginners/mydict"
)

func main() {
	dictionary := mydict.Dictioinary{"first": "First word"}
	word := "hello"
	definition := "Greeting"
	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}

	hello, _ := dictionary.Search(word)
	fmt.Println("found", word, "definition: ", hello)

	err2 := dictionary.Add(word, definition)
	if err2 != nil {
		fmt.Println(err2)
	}
}
