package main

import "fmt"

const chinese = "chinese"
const french = "french"
const spanish = "spanish"

const englishPrefix = "Hello, "
const chinesePrefix = "你好, "
const frenchPrefix = "Bonjour, "
const spanishPrefix = "Hola, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return getPrefix(language) + name

}

func getPrefix(language string) (prefix string) {
	switch language {
	case chinese:
		prefix = chinesePrefix
	case french:
		prefix = frenchPrefix
	case spanish:
		prefix = spanishPrefix
	default:
		prefix = englishPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("", "sss"))
}
