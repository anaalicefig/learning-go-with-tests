package main

import "fmt"

const helloPrefixEnglish = "Hello "

func Hello(name string, idioma string) string {
	if name == "" {
		name = "World"
	}

	if idioma == "spanish" {
		return "Hola, " + name
	}

	return helloPrefixEnglish + name
}

func main() {
	fmt.Println(Hello("mundo", "spanish"))
}
