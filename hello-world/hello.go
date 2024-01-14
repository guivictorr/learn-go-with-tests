package main

import "fmt"

const (
	spanish            = "Spanish"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	defaultName        = "World"
)

func Hello(name, language string) string {
	if name == "" {
		name = defaultName
	}

	if language == spanish {
		return spanishHelloPrefix + name
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("World", "English"))
}
