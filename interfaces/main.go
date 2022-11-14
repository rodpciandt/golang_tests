package main

import "fmt"

type bot interface {
	getGreeting() string // if some type has the getGreeting method and retuns a string, it will be recognized as bot
}

type englishBot struct {}
type spanishBot struct {}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
	
}



func (englishBot) getGreeting() string {
	// very custom logic
	return "Hi, there!"
}
// you can ommit the value if you are not going to use it
func (spanishBot) getGreeting() string {
	return "Hola"
}