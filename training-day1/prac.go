package main

import "fmt"

var Country = "USA"

func prac() {
	fmt.Println(Country)

	Country := "India"

	fmt.Println(Country)

	x, y := 10, 20
	x, y = y, x

	fmt.Println(x, y)

	fn := func(num int) {
		fmt.Println("Hello", num)
	}

	test(fn)

	var w [5]float64
	fmt.Println(w)

	sl := make([]int, 3, 5)
	sl[0] = 10
	sl[1] = 20
	sl[2] = 30
	sl = append(sl, 40)

	fmt.Println(sl)

	key := "key1"
	mp := map[string]string{
		key: "Hello map",
	}

	fmt.Println(mp[key])

	fn2 := func() (int, bool) {
		return 10, false
	}

	if num, exist := fn2(); exist {
		fmt.Println(num)
	} else {
		fmt.Println("Not number")
	}
}

func test(fn func(int)) {
	fn(10)
}
