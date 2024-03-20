package itunes_library_parser

import (
	"fmt"
	"os"
	"strings"
)

type Value struct {
	str string
	typ string
}

func (v Value) IsString() bool {
	return v.typ == "string"
}

func (v Value) IsInt() bool {
	return v.typ == "integer"
}

func (v Value) IsDate() bool {
	return v.typ == "date"
}

func ParseSong(data string) (map[string]Value, error) {
	songDetails := make(map[string]Value)

	for _, line := range strings.Split(data, "\n") {
		// Field name
		line = strings.TrimPrefix(line, "<key>")
		end := strings.Index(line, "</key>")
		fieldName := line[:end]

		// Value type
		line = line[end+6:]
		end = strings.IndexByte(line, '>')
		valueType := line[:end]

		// Value string
		line = line[end+1:]
		end = strings.IndexByte(line, '<')
		value := line[:end]

		songDetails[fieldName] = Value{
			str: value,
			typ: valueType,
		}
	}

	return songDetails, nil
}

func main() {
	file, err := os.Open("Bibliothèque.xml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

}
