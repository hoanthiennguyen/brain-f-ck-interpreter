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

	result := []*Instruction{}
	var prevOp *Instruction
	for index := 0; index < len(srcCode); index++ {
		e := srcCode[index]
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
		case ']':
			in = NewInstruction(OpEndLoop)
		}

		if index >= 1 {
			if prevOp.Op == in.Op && prevOp.Op.CanStack() {
				prevOp.Param++

				// if this is the last operation, add to the result
				if index == len(srcCode)-1 {
					result = append(result, prevOp)
				}
			} else {
				result = append(result, prevOp)
				prevOp = in
				// last operation need to be added
				if index == len(srcCode)-1 {
					result = append(result, in)
				}
			}
		} else {
			prevOp = in
		}

	}
	buildMatchingBrackets(result)

	return result
}

func buildMatchingBrackets(ins []*Instruction) {
	stk := NewStack[*Instruction]()
	for index, in := range ins {
		switch in.Op {
		case OpBeginLoop:
			in.data = index
			stk.Push(in)
		case OpEndLoop:
			begin, _ := stk.Pop()
			begin.Param = index
			in.Param = begin.data
		}
	}
}
