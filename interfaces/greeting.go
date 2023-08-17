package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func (englishBot) getGreeting() string {
	return "Hello World!!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

// without interfaces
func printEnglishMesage(eb englishBot) {
	msg := eb.getGreeting()
	fmt.Println(msg)
}

func printSpanishMessage(sb spanishBot) {
	msg := sb.getGreeting()
	fmt.Println(msg)
}

// with interfaces
type bot interface {
	getGreeting() string
}

func printMessage(b bot) {
	msg := b.getGreeting()
	fmt.Println(msg)
}
