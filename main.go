package main

import (
	"fmt"
	"os"
)

func greeter(name string) string {
	return fmt.Sprintf("Hello %s!", name)
}

func main() {
	var noun string
	if len(os.Args) < 2 {
		noun = "World"
	} else {
		noun = os.Args[1]
	}
	fmt.Println(greeter(noun))
}
