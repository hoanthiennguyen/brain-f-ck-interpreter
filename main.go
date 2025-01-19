package main

import (
	"fmt"
	"os"

	"github.com/hoanthiennguyen/go-stack"
)

func execute(srcCode string) {
	memory := make([]int, 30000)
	srcCounter := 0
	dataPointer := 0
	// Store parathensis position
	stk := stack.New[int]()
	var err error
	for srcCounter < len(srcCode) {
		instruction := rune(srcCode[srcCounter])
		switch instruction {
		case '+':
			memory[dataPointer] += 1
			srcCounter++
		case '-':
			memory[dataPointer] -= 1
			srcCounter++
		case '>':
			dataPointer++
			srcCounter++
		case '<':
			if dataPointer > 0 {
				dataPointer--
			}
			srcCounter++
		case '.':
			data := rune(memory[dataPointer])
			fmt.Print(string(data))
			srcCounter++
		case ',':
			input := int('i')
			memory[dataPointer] = input
			srcCounter++
			// TODO
		case '[':
			if memory[dataPointer] > 0 {
				stk.Push(srcCounter)
				srcCounter++
			} else {
				for srcCode[srcCounter] != ']' {
					srcCounter++
				}
				srcCounter++
			}
		case ']':
			if stk.IsEmpty() {
				panic("Stack is empty")
			}
			srcCounter, err = stk.Pop()
			if err != nil {
				panic(err)
			}
		default:
			srcCounter++
		}
	}
}

func main() {
	fileName := os.Args[1]
	contentRaw, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	execute(string(contentRaw))
}

// func isValidChar(c rune) bool {
// 	return slices.Contains([]rune{'+', '-', '>', '<', '.', '[', ']', ','}, c)
// }
