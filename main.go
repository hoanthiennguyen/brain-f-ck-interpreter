package main

import (
	brainfuck "brain-fuck/src"
	"os"
)

func main() {
	f1 := os.Args[1]
	switch f1 {
	case "compile":
		brainfuck.CompileV2(os.Args[2])
	case "run":
		brainfuck.Run(os.Args[2])
	}
}
