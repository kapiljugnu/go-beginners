package main

import "fmt"

func evenOddCheck() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, num := range numbers {
		isEven := num%2 == 0
		if isEven {
			fmt.Println(num, "is Even")
		} else {
			fmt.Println(num, "is Odd")
		}
	}
}
