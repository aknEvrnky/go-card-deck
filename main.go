package main

import "fmt"

func main() {
	d := fromFile("data.txt")

	fmt.Println(d.toString())
	// cards.saveToFile("data.txt")
}
