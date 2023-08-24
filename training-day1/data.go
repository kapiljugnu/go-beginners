package main

import "fmt"

var data map[string]string

func init() {
	data = make(map[string]string)
}

func addItem(k, v string) {
	fmt.Println("=======addItem======")
	if _, ok := data[k]; ok {
		fmt.Println("Key already exist")
		return
	}
	data[k] = v
	fmt.Println("New record has been inserted")
}

func updateItem(k, v string) {
	fmt.Println("=======updateItem======")
	if _, ok := data[k]; !ok {
		fmt.Println(k, "not found")
		return
	}
	data[k] = v
	fmt.Println("Item updated successfully")
}

func getById(k string) {
	fmt.Println("=======getById======")
	value, ok := data[k]
	if !ok {
		fmt.Println(k, "not found")
		return
	}

	fmt.Println(k, "value is", value)
}

func getAll() {
	fmt.Println("=======getAll======")
	if len(data) == 0 {
		fmt.Println("No data found.")
		return
	}

	for k, v := range data {
		fmt.Println(k, "value is", v)
	}
}

func deleteItem(k string) {
	fmt.Println("=======deleteItem======")
	if _, ok := data[k]; !ok {
		fmt.Println(k, "not found")
		return
	}
	delete(data, k)
	fmt.Println("Item delete successfully")
}
