package main

import (
	"fmt"
	"os"
)

func main() {
	command := os.Args[1]

	switch command {
	case "add":
		fmt.Println("Add command")

	case "update":
		fmt.Println("Update command")

	case "delete":
		fmt.Println("delete command")

	case "list":
		if os.Args[2] == "done" {
			fmt.Println("list done")
		} else if os.Args[2] == "todo" {
			fmt.Println("List todo")
		} else if os.Args[2] == "in-progress" {
			fmt.Println("List in-progress")
		} else {
			fmt.Println("Wrong command. Check tutotial")
		}

	case "mark-in-progress":
		fmt.Println("Mark-in-progress command")

	case "mark-done":
		fmt.Println("Mark-done command")

	default:
		fmt.Println("Wrong command")
	}
}
