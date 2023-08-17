package main

import (
	"fmt"
	"io"
	"os"
)

func readFile(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, f)
}
