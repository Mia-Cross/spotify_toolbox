package itunes_library_parser

import (
	"fmt"
	"os"
	"strings"
)

type Library struct {
	Songs     []Song
	Playlists []Playlist
}

func parseLibrary(data string) (Library, error) {
	library := Library{}

	_ = strings.Split(data, "\n")

	return library, nil
}

func main() {
	file, err := os.ReadFile("Bibliothèque.xml")
	if err != nil {
		fmt.Printf("error open file: %s", err)
		os.Exit(1)
	}

	_, err = parseLibrary(string(file))
	if err != nil {
		fmt.Printf("error parsing library: %s", err)
		os.Exit(1)
	}

}
