package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to Go REPL. Type 'exit()' to quit.")

	reader := bufio.NewReader(os.Stdin)
	lines := []string{}
	fmt.Print(">> ") // Prompt
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		if input == "exit()" {
			break
		}

		lines = append(lines, input)
		if strings.HasSuffix(input, ";") {
			output := strings.Join(lines, "\n")
			fmt.Println(output)
			lines = []string{}
			fmt.Print(">> ") // Prompt
		} else {
			// continue to read input
			fmt.Print("> ")
		}

	}
}
