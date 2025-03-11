package main

import (
	"fmt"
	"os"
	"slices"

	"github.com/hoanthiennguyen/go-stack"
)

func execute(srcCode string) {
	memory := make([]int, 1000000)
	srcCounter := 0
	dataPointer := 0
	// Store parathensis position
	stk := stack.New[int]()
	for srcCounter < len(srcCode) {
		instruction := rune(srcCode[srcCounter])
		switch instruction {
		case '+':
			memory[dataPointer]++
		case '-':
			memory[dataPointer]--
		case '>':
			dataPointer++
		case '<':
			if dataPointer > 0 {
				dataPointer--
			} else {
				panic("invalid <")
			}
		case '.':
			data := rune(memory[dataPointer])
			fmt.Print(string(data))
		case ',':
			input := int('i')
			memory[dataPointer] = input
			// TODO
		case '[':
			if memory[dataPointer] > 0 {
				stk.Push(srcCounter)
			} else {
				dept := 1
				for dept > 0 {
					srcCounter++
					if srcCounter == len(srcCode) {
						panic("Unmatched [")
					}

					if srcCode[srcCounter] == ']' {
						dept--
					}
					if srcCode[srcCounter] == '[' {
						dept++
					}
				}
			}
		case ']':
			if memory[dataPointer] > 0 {
				lastMatching, err := stk.Peek()
				if err != nil {
					panic(err)
				}
				srcCounter = lastMatching
			} else {
				_, err := stk.Pop()
				if err != nil {
					panic(err)
				}
			}
		}
		srcCounter++

	}
}

func filterComment(srcCode string) string {
	result := []rune{}
	validChars := []rune{'+', '-', '>', '<', '.', '[', ']', ','}
	for _, c := range srcCode {
		if slices.Contains(validChars, c) {
			result = append(result, c)
		}
	}

	return string(result)
}

func main() {
	fileName := os.Args[1]
	contentRaw, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	content := filterComment(string(contentRaw))
	execute(content)
}

//	func isValidChar(c rune) bool {
//		return slices.Contains([]rune{'+', '-', '>', '<', '.', '[', ']', ','}, c)
//	}
