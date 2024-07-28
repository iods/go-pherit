package greeting

const defaultPrefix = "Hello, "
const defaultSuffix = "World!"

var greetingPrefixLanguage = map[string]string{
	"English": "Hello, ",
	"Spanish": "Hola, ",
	"French":  "Bonjour, ",
	"German":  "Hallo, ",
	"Italian": "Ciao, ",
}

func setGreetingPrefix(language string) string {
	greetingPrefix := greetingPrefixLanguage[language]
	if greetingPrefix == "" {
		greetingPrefix = defaultPrefix
	}
	return greetingPrefix
}

func setName(name string) string {
	if name == "" {
		name = defaultSuffix
	}
	return name
}

func Greet(name string, language string) string {
	name = setName(name)
	greetingPrefix := setGreetingPrefix(language)
	return greetingPrefix + name
}
