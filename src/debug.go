package brainfuck

import (
	"encoding/json"
	"fmt"
)

func printDebug(prefix string, data any) {
	jsonRaw, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s: %s\n", prefix, string(jsonRaw))
}
