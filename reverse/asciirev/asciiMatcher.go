package asciirev

import (
	"fmt"
	"log"
)

func AsciiMatcher(mapGraph map[int]string, data []string) string {
	slice := []rune{}

	str := ""

	mapData := RevMapper(data)

	for i := 1; i <= len(mapGraph); i++ {
		for key, data := range mapData {
			if mapGraph[i] == data {
				slice = append(slice, key)
			}
		}
	}
	if len(slice) != len(mapGraph) {
		error := "Error: no match found, please check if the symbols are correctly formated"
		fmt.Println("After each graphic letter there must be a space before next one")
		log.Fatal(error)
	}

	slice = append(slice, 10)

	for _, val := range slice {
		str += string(val)
	}

	return str
}
