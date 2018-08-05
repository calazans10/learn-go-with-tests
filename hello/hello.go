package hello

const french = "French"
const portuguese = "Portuguese"
const spanish = "Spanish"
const englishPrefix = "Hello, "
const frenchHelloPrefix = "Bonjour, "
const portugueseHelloPrefix = "Olá, "
const spanishHelloPrefix = "Hola, "

// Hello returns a greeting message
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name

}

func greetingPrefix(language string) string {
	switch language {
	case french:
		return frenchHelloPrefix
	case portuguese:
		return portugueseHelloPrefix
	case spanish:
		return spanishHelloPrefix
	default:
		return englishPrefix
	}
}
