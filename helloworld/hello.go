package main

import "fmt"

const (
	spanish       = "Spanish"
	french        = "French"
	englishPrefix = "Hello, "
	spanishPrefix = "Hola, "
	frenchPrefix  = "Bonjour, "
)

func Hello(name string, language string) string {
	prefix := greetingPrefix(language)
	to := name
	if len(name) == 0 {
		to = "peeps"
	}

	return prefix + to
}

// prefix is named return value
func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	default:
		prefix = englishPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Jack", ""))
}
