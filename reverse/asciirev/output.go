package asciirev

import (
	"fmt"
	"os"
)

func Output(filename, inputStr, banner string) {
	data, err := FileReader(banner)
	if err != nil {
		// log.Fatalln(err)
		fmt.Println(err.Error())
		return
	}
	mapData := AsciiMapper(data)
	output := AsciiPrinter(mapData, inputStr)
	if len(filename) != 0 {
		output = output + "\n" + "\n"
		os.WriteFile(filename, []byte(output), 0o644)

	}
}
