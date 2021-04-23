package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("What manual page do you want?")
		fmt.Println("For example, try 'oc man index'")
	} else {
		target := os.Args[1]
		if target == "index" {
			fmt.Println("Printing index...")
		} else {
			fmt.Println("I don't have a page for", target)
		}
	}
}
