package brainfuck

import (
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

func Compile(fileName string) {
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
