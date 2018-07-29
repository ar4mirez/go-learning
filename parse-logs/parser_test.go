package logs_test

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	logs "github.com/ar4mirez/go-learning/parse-logs"
)

func TestParse(t *testing.T) {

	tt := []struct {
		name     string
		filename string
		result   map[string]int
	}{
		{
			name:     "log_01",
			filename: "log_01.txt",
			result: map[string]int{
				"google.com": 1,
				"bing.com":   2,
				"yahoo.com":  3,
				"ask.com":    2,
			},
		},
		{
			name:     "log_02",
			filename: "log_02.txt",
			result: map[string]int{
				"google.com": 4,
				"bing.com":   1,
				"yahoo.com":  2,
				"ask.com":    1,
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			err := logs.Parse(tc.filename)
			if err != nil {
				t.Errorf("failed parsing: %v", err)
			}

			f, err := os.Open(fmt.Sprintf("records_%s", tc.filename))
			if err != nil {
				t.Errorf("unable to open file: %v", err)
			}

			scanner := bufio.NewScanner(f)

			for scanner.Scan() {
				data := strings.Split(scanner.Text(), " ")
				key := fmt.Sprintf("%s", data[0])
				val := tc.result[key]
				v, _ := strconv.Atoi(data[1])
				if val != v {
					t.Errorf("expected %s %d and got %d", key, val, v)
				}
			}
		})
	}
}
