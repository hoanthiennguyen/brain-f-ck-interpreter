package brainfuck

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
	"slices"
	"strings"
)

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

func findMatchBrackets(srcCode string) [][2]int {
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

func CompileV2(fileName string) {
	contentRaw, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	instructions := compileV2(string(contentRaw))

	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	if err := encoder.Encode(instructions); err != nil {
		panic(err)
	}

	newFile := getCompiledFileName(fileName)
	if err := os.WriteFile(
		newFile,
		buffer.Bytes(),
		os.ModePerm,
	); err != nil {
		panic(err)
	}

	fmt.Printf("Compiled successfully %s\n", newFile)
}

func getCompiledFileName(fileName string) string {
	nameArr := strings.Split(fileName, ".")
	fileNameCompiled := nameArr[0] + ".bfc"
	return fileNameCompiled
}

func compileV2(srcCode string) []*Instruction {
	srcCode = filterComment(string(srcCode))
	matchBrackets := findMatchBrackets(srcCode)
	openMap, closeMap := make(map[int]int), make(map[int]int)
	for _, e := range matchBrackets {
		openB, closeB := e[0], e[1]
		openMap[closeB] = openB
		closeMap[openB] = closeB
	}

	result := []*Instruction{}
	for index, e := range srcCode {
		var in *Instruction
		switch e {
		case '+':
			in = NewInstruction(OpIncr)
		case '-':
			in = NewInstruction(OpDecr)
		case '>':
			in = NewInstruction(OpMoveRight)
		case '<':
			in = NewInstruction(OpMoveLeft)
		case '.':
			in = NewInstruction(OpOutput)
		case ',':
			in = NewInstruction(OpInput)
		case '[':
			in = NewInstruction(OpBeginLoop)
			in.Param = closeMap[index]
		case ']':
			in = NewInstruction(OpEndLoop)
			in.Param = openMap[index]
		}

		result = append(result, in)
	}

	return result
}
