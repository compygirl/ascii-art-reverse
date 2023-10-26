package asciirev

func AsciiMapper(data []string) map[rune][]string {
	mapData := map[rune][]string{}
	runes := 32
	for len(data) != 0 {
		elements := data[1:9]
		mapData[rune(runes)] = elements
		data = data[9:]
		runes++
	}
	return mapData
}
