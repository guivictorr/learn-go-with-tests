package main

import "fmt"

const (
	spanish            = "Spanish"
	french             = "French"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	defaultName        = "World"
)

func Hello(name, language string) string {
	if name == "" {
		name = defaultName
	}

	if language == french {
		return frenchHelloPrefix + name
	}

	if language == spanish {
		return spanishHelloPrefix + name
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("World", "English"))
}
