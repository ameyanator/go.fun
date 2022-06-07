package main

const englishHelloPrefix = "Hello "
const hindiHelloPrefix = "Namaste "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	if language == "English" {
		return englishHelloPrefix + name + "!"
	}
	return hindiHelloPrefix + name + "!"
}

func main() {
	// fmt.Println(Hello("World"))
}
