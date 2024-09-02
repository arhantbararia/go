package main

import "fmt"

const (
	englishHello = "Hello, "
	spanishhello = "Hola, "
	frenchhello  = "Bon, "
)

func Hello(name string, language string) string {

	if name == "" {
		return greetingPrefix(language) + "World!"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "French":
		prefix = frenchhello

	case "Spanish":
		prefix = spanishhello

	default:
		prefix = englishHello
	}

	return
}

//to write and test
// 1. hello world using parameter and if empty name then print Hello
// 2. hello in other languages using switch
// 3.

func main() {
	fmt.Println(Hello("Arhant Main", ""))
}
