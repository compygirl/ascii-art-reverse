package asciirev

func RevMapper(data []string) map[rune]string {
	mapData := map[rune]string{}

	runes := 32
	for len(data) != 0 {
		elements := data[1:9]
		str := ""
		for _, val := range elements {
			str += val + "\n"
		}
		mapData[rune(runes)] = str
		data = data[9:]

		runes++

	}
	return mapData
}
