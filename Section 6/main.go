package Section_6

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func main() {

}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}

func (englishBot) getGreeting() string {
	// Very Custom logic for generating and english greeting
	// if we are not using the receiver variable we can omit it from (eb englishbot)
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
