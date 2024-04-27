package main

import "fmt"

const (
	french             = "French"
	spanish            = "Spanish"
	englishSuffixHello = "Hello, "
	frenchSuffixHello  = "Bonjour, "
	spanishSuffixHello = "Hola, "
)

func SayHello(name, lang string) string {
	if name == "" {
		name = "world"
	}
	return prefixGreetings(lang) + name
}
func prefixGreetings(lang string) string {
	switch lang {
	case french:
		return frenchSuffixHello
	case spanish:
		return spanishSuffixHello
	default:
		return englishSuffixHello
	}
}

func main() {
	fmt.Println(SayHello("", ""))
}
