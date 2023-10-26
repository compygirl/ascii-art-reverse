package asciirev

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func FileReader(filename string) ([]string, error) {
	if !strings.HasSuffix(filename, ".txt") {
		filename += ".txt"
	}
	var data []string

	file, err := os.Open(filename)
	if err != nil {
		return data, fmt.Errorf("open file err: %v", err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	defer file.Close()

	return data, nil
}
