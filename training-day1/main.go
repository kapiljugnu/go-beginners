package main

func main() {
	addItem("firstname", "Kapil")
	addItem("firstname", "John")
	addItem("lastname", "Kumar")
	updateItem("lastname", "Jugnu")
	updateItem("nokey", "novalue")
	getById("firstname")
	getAll()
	deleteItem("nokey")
	deleteItem("lastname")
	getAll()
}
