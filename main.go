package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/go-risc-v/converter"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter value: ")
		ok := scanner.Scan()
		if !ok {
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "Error reading input:", err)
			}
			break
		}
		line := scanner.Text()
		result := converter.Convert(line)
		fmt.Println("Converted value:", result)
	}
}
