package main

import (
	"fmt"

	"github.com/Hwisaek/Go/nomadcoder_go-for-beginners/mydict"
)

func main() {
	dictionary := mydict.Dictioinary{"first": "First word"}
	definition, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
}
