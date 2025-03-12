package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/hoanthiennguyen/go-stack"
)

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
	f1 := os.Args[1]
	switch f1 {
	case "compile":
		compile(os.Args[2])
	case "run":
		run(os.Args[2])
	}
}

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

func buildClosingBracketMap(srcCode string) [][2]int {
	closingBracket := [][2]int{}
	for i := 0; i < len(srcCode); i++ {
		instruction := rune(srcCode[i])
		if instruction == '[' {
			tmp := i + 1
			countMissingBracket := 1
			for tmp < len(srcCode) {
				if rune(srcCode[tmp]) == ']' {
					countMissingBracket--
					if countMissingBracket == 0 {
						break
					}

				} else if rune(srcCode[tmp]) == '[' {
					countMissingBracket++
				}

				tmp++
			}

			closingBracket = append(closingBracket, [2]int{i, tmp})
		}
	}

	return closingBracket
}

func compile(fileName string) {
	contentRaw, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	srcCode := filterComment(string(contentRaw))

	closingBracketsMap := buildClosingBracketMap(srcCode)
	arr := []string{}
	for _, e := range closingBracketsMap {
		arr = append(arr, fmt.Sprint(e[0]), fmt.Sprint(e[1]))
	}

	serializeString := strings.Join(arr, ",")
	serializeString = serializeString + "\n" + srcCode

	nameArr := strings.Split(fileName, ".")
	fileNameCompiled := nameArr[0] + ".bfc"
	if err := os.WriteFile(fileNameCompiled, []byte(serializeString), os.ModePerm); err != nil {
		panic(err)
	}
	fmt.Printf("Compiled successfully %s\n", fileNameCompiled)
}

//	func isValidChar(c rune) bool {
//		return slices.Contains([]rune{'+', '-', '>', '<', '.', '[', ']', ','}, c)
//	}
