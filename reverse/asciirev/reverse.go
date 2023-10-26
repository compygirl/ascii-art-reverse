package asciirev

import (
	"fmt"
)

func CheckFormat(graph []string) string {
	for _, str := range graph {
		for ind, _ := range str {
			if str[ind] != 32 && (str[ind] == 'O' || str[ind] == 'o' || str[ind] == '0') {
				return "thinkertoy"
			}
		}
	}

	for _, str := range graph {
		for ind, _ := range str {
			if str[ind] != 32 && (str[ind] == '\\' || str[ind] == '/' || str[ind] == '(' || str[ind] == ')' || str[ind] == 'V') {
				return "standard"
			}
		}
	}
	return "shadow"
}

func Reverse(filename string) {
	graph, err2 := FileReader(filename)
	if err2 != nil {
		// log.Fatalln(err2)
		fmt.Println(err2.Error())
		return
	}

	fontformat := CheckFormat(graph)
	data, err1 := FileReader(fontformat)
	if err1 != nil {
		// log.Fatalln(err1)
		fmt.Println(err1.Error())
		return
	}
	GraphicMapper(graph, data)
}
