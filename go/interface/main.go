package main

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}
func printGreeting(b bot) {
	println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	//very custom logic for generating an english greeting
	return "Hi There!"
}
func (spanishBot) getGreeting() string {
	//very custom logic for generating an spanish greeting
	return "Hola!"
}
