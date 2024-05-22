package main

import (
	"bufio"
	"fmt"
	"github.com/samorokowsky/lessongo/greet"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	name := getCLIValue(reader, "name")
	tableStr := getCLIValue(reader, "table")

	table, err := strconv.Atoi(tableStr)

	if err != nil {
		fmt.Printf("Error during table number conversion, got %s\n", tableStr)
		return
	}

	text := greet.Greet(name, table)
	fmt.Println(text)
}

func getCLIValue(reader *bufio.Reader, propertyName string) string {
	fmt.Printf("Enter %s: ", propertyName)

	readValue, _ := reader.ReadString('\n')
	readValue = strings.Trim(readValue, "\r\n")

	return readValue
}
