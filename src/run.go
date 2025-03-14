package brainfuck

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
)

func Run(compiledFileName string) {
	contentRaw, err := os.ReadFile(compiledFileName)
	if err != nil {
		panic(err)
	}

	buffer := *bytes.NewBuffer(contentRaw)
	decoder := gob.NewDecoder(&buffer)

	instructions := []*Instruction{}
	if err := decoder.Decode(&instructions); err != nil {
		panic(err)
	}

	execute(instructions)
}

func execute(arr []*Instruction) {
	memory := make([]int, 100000)
	srcCounter := 0
	dataPointer := 0
	// Store parathensis position
	for srcCounter < len(arr) {
		instruction := arr[srcCounter]
		switch instruction.Op {
		case OpIncr:
			memory[dataPointer]++
		case OpDecr:
			memory[dataPointer]--
		case OpMoveRight:
			dataPointer++
		case OpMoveLeft:
			dataPointer--
		case OpOutput:
			data := rune(memory[dataPointer])
			fmt.Print(string(data))
		case OpInput:
			input := int('i')
			memory[dataPointer] = input
			// TODO
		case OpBeginLoop:
			if memory[dataPointer] == 0 {
				srcCounter = instruction.Param
			}
		case OpEndLoop:
			if memory[dataPointer] > 0 {
				srcCounter = instruction.Param
			}
		}
		srcCounter++

	}
}
