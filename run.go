package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/hoanthiennguyen/go-stack"
)

func run(compiledFileName string) {
	contentRaw, err := os.ReadFile(compiledFileName)
	if err != nil {
		panic(err)
	}

	arr := strings.Split(string(contentRaw), "\n")
	firstLine := arr[0]
	srcCode := arr[1]

	arr = strings.Split(firstLine, ",")
	closingBracket := make(map[int]int)
	for index := 0; index < len(arr); index += 2 {
		open, _ := strconv.Atoi(arr[index])
		closing, _ := strconv.Atoi(arr[index+1])
		closingBracket[open] = closing
	}

	execute(srcCode, closingBracket)
}

func execute(srcCode string, correspondingClosingBracket map[int]int) {
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
				srcCounter = correspondingClosingBracket[srcCounter]
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
