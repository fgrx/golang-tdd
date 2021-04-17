package main

import "fmt"

const englishPrefix="Hello"
const frenchPrefix="Bonjour"
const spanishPrefix="Hola"

func Hello(name string,lang string) string{
	if name==""{
		name="world"
	}
	prefix:=getPrefix(lang)
	return prefix+" "+name
}

func getPrefix(lang string)string{
	switch lang{
		case "fr":
			return frenchPrefix
		case "es":
			return spanishPrefix
		default:
			return englishPrefix
	}
}

func main() {
	fmt.Println(Hello("Fabien",""),"")
}