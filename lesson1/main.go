package main

import (
	"fmt"
	"github.com/samorokowsky/lessongo/greet"
	"os"
	"strconv"
)

func main() {
	name := os.Args[1]
	tableStr := os.Args[2]

	table, err := strconv.Atoi(tableStr)

	if err != nil {
		fmt.Printf("Error during table number conversion, got %s\n", tableStr)
		return
	}

	text := greet.Greet(name, table)
	fmt.Println(text)
}
