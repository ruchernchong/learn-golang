package main

import "fmt"

func main() {
	var (
		title, description = "Learn Golang", "Learning the basics of Golang with me"
	)

	var chapter int = 1
	var isCompleted = true

	fmt.Println("Title: \n", title)
	fmt.Println("Description: \n", description)
	fmt.Println("Chapter: \n", chapter)
	fmt.Println("Lesson completed? \n", isCompleted)
}
