package hello

import "fmt"

const englishHelloPrefix = "Hello, "
const portugueseHelloPrefix = "Olá, "
const frenchHelloPrefix = "Bonjour, "
const portuguese = "Portuguese"
const french = "French"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case portuguese:
		prefix = portugueseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
