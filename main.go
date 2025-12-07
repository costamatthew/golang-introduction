package main

import (
	"github.com/costamatthew/golang-introduction/meet"
)

func main() {
	var texto1 string = "Texto01"
	var texto2 = "Texto02"
	texto3 := "Texto03"

	const texto4 string = "Texto04"
	const texto5 = "Texto05"

	meet.SayHello()
	meet.Say("Hello")
	meet.Say(texto1)
	meet.Say(texto2)
	meet.Say(texto3)
	meet.Say(texto4)
	meet.Say(texto5)
}
