package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printEnglishMesage(eb)
	printSpanishMessage(sb)

	printMessage(eb)
	printMessage(sb)

	// second project which make http request
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// commented for below implementation
	// fmt.Println(resp)

	// Reader interface to understand interface more
	// bs := []byte{} // instead use make

	bs := make([]byte, 99999) // "make" is in built func that take type and the size of empty space with which we want slice to intialize with
	resp.Body.Read(bs)
	// // fmt.Println(string(bs))

	// one liner for code line 32 to 34
	// with writer interface
	// io.Copy(os.Stdout, resp.Body)

	t := triangle{base: 10, height: 10}
	s := square{sideLength: 10}

	printArea(t)
	printArea(s)

	fmt.Println("-------------------------")
	filename := os.Args[1]
	readFile(filename)
}
