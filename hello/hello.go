package hello

const (
	englishPrefix = "Hello, "
	spanishPrefix = "Hola, "
	frenchPrefix  = "Salute, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishPrefix

	switch language {
	case "French":
		prefix = frenchPrefix
	case "Spanish":
		prefix = spanishPrefix
	}
	return prefix + name
}