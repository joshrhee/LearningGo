package main

import (
	"fmt"

	"github.com/joshrhee/dict/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "Frist word"}
	word := "hello"
	definition := "greeting"
	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}
	hello, err := dictionary.Search(word)
	fmt.Println("found: ", word, " definition: ", hello)

	err2 := dictionary.Add(word, definition)
	if err2 != nil {
		fmt.Println(err2)
	}

	err3 := dictionary.Update(word, "second")
	if err != nil {
		fmt.Println(err3)
	}
	word, _ = dictionary.Search(word)
	fmt.Println("Updated word is: ", word)

	dictionary.Delete(word)
	word, err4 := dictionary.Search(word)
	if err4 != nil {
		fmt.Println(err4)
	}
	fmt.Println("After delete, word is: ", word)
}
