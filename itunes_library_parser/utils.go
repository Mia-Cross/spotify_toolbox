package itunes_library_parser

import (
	"errors"
	"strconv"
	"strings"
)

func readKey(line string) (string, int) {
	line = strings.TrimPrefix(line, "<key>")
	end := strings.Index(line, "</key>")
	return line[:end], end
}

func parseLine(line string) (key, value string) {
	line = strings.TrimSpace(line)

	// Field name
	key, end := readKey(line)

	// Value type
	line = line[end+6:]
	end = strings.IndexByte(line, '>')
	//valueType := line[:end]

	// Value string
	line = line[end+1:]
	end = strings.IndexByte(line, '<')
	if end == -1 {
		return key, ""
	}
	value = line[:end]

	return key, value
}

func parseArray(array []string) []int {
	trackIDs := []int(nil)

	for _, line := range array {
		line = strings.TrimSpace(line)
		if !strings.HasPrefix(line, "<key>") {
			continue
		}
		_, value := parseLine(line)
		trackID, _ := strconv.Atoi(value)
		trackIDs = append(trackIDs, trackID)
	}

	return trackIDs
}

func getDictLimit(data []string) (int, error) {
	endLine := "</dict>"
	for i := 0; data[0][i] == '\t'; i++ {
		endLine = "\t" + endLine
	}

	for i, line := range data {
		if line == endLine {
			return i, nil
		}
	}
	return 0, errors.New("couldn't find end of dict")
}
