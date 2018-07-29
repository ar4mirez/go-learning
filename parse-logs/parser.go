package logs

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Parse a log file and counts the requests per host.
func Parse(filename string) error {

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	records := make(map[string]int, 200000)

	if err := scanner.Err(); err != nil {
		return err
	}

	for scanner.Scan() {
		data := strings.Split(scanner.Text(), " ")
		records[fmt.Sprintf("%s", data[0])]++
	}

	result, err := os.Create(fmt.Sprintf("records_%s", filename))
	if err != nil {
		return err
	}
	defer result.Close()

	for k, v := range records {
		result.WriteString(fmt.Sprintf("%s %d\n", k, v))
	}

	return nil
}
