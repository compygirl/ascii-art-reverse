package asciirev

import (
	"strings"
)

func AsciiPrinter(mapData map[rune][]string, inputStr string) string {
	inputStr = strings.ReplaceAll(inputStr, "\n", "\\n")
	inputSlice := strings.Split(inputStr, "\\n")
	output := ""
	for _, word := range inputSlice {
		if word != "" {
			for ind := 0; ind < 8; ind++ {
				for _, character := range word {
					output += mapData[character][ind]
				}
				output += "\n"
			}
		} else {
			output += "\n"
		}
	}
	output = output[:len(output)-1]
	return output
}
