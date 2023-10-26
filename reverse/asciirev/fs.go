package asciirev

import (
	"fmt"
	// "log"
)

func FS(inputStr, banner string) {
	data, err := FileReader(banner)
	if err != nil {
		// log.Fatalln(err)
		fmt.Println(err.Error())
		return
	}
	mapData := AsciiMapper(data)
	output := AsciiPrinter(mapData, inputStr)
	fmt.Println(output)
}
