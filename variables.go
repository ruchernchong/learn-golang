package main

import "fmt"

func main() {
	lesson()
}

func lesson() {
	title := "Learn Golang"
	description := "Learning the basics of Golang with me"
	chapter := 1

	fmt.Println("Title: \n", title)
	fmt.Println("Description: \n", description)
	fmt.Println("Chapter: \n", chapter)
}
