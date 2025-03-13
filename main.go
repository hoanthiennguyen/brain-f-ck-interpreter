package main

import (
	"os"
)

func main() {
	f1 := os.Args[1]
	switch f1 {
	case "compile":
		compile(os.Args[2])
	case "run":
		run(os.Args[2])
	}
}
