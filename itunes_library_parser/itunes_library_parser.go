package itunes_library_parser

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Song struct {
	TrackID int

	Name        string
	Artist      string
	AlbumArtist string
	Album       string

	Kind      string
	Size      int
	TotalTime int
	DateAdded time.Time
	PlayCount int

	SortAlbum       string
	SortAlbumArtist string

	Location string
}

func parseLine(line string) (key, value string) {
	line = strings.TrimSpace(line)

	// Field name
	line = strings.TrimPrefix(line, "<key>")
	end := strings.Index(line, "</key>")
	key = line[:end]

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

func ParseSong(data string) (Song, error) {
	song := Song{}
	var err error

	for _, line := range strings.Split(data, "\n") {
		if len(line) == 0 {
			continue
		}
		key, value := parseLine(line)

		switch key {
		case "Track ID":
			song.TrackID, err = strconv.Atoi(value)
			if err != nil {
				return song, err
			}
		case "Name":
			song.Name = value
		case "Artist":
			song.Artist = value
		case "Album Artist":
			song.AlbumArtist = value
		case "Album":
			song.Album = value
		case "Kind":
			song.Kind = value
		case "Size":
			song.Size, err = strconv.Atoi(value)
			if err != nil {
				return song, err
			}
		case "Total Time":
			song.TotalTime, err = strconv.Atoi(value)
			if err != nil {
				return song, err
			}
		case "Date Added":
			song.DateAdded, err = time.Parse(time.RFC3339, value)
			if err != nil {
				return song, err
			}
		case "Play Count":
			song.PlayCount, err = strconv.Atoi(value)
			if err != nil {
				return song, err
			}
		case "Sort Album":
			song.SortAlbum = value
		case "Sort Album Artist":
			song.SortAlbumArtist = value
		case "Location":
			song.Location = value
		}
	}

	return song, nil
}

type Playlist struct {
	Name        string
	Description string
	PlaylistID  int
	TrackIDs    []int
}

func parsePlaylistHeader(line string, playlist *Playlist) error {
	var err error
	key, value := parseLine(line)

	switch key {
	case "Name":
		playlist.Name = value
	case "Description":
		playlist.Description = value
	case "Playlist ID":
		playlist.PlaylistID, err = strconv.Atoi(value)
		if err != nil {
			return err
		}
	}
	return nil
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

func ParsePlaylist(data string) (Playlist, error) {
	playlist := Playlist{}
	lines := strings.Split(data, "\n")

	for i, line := range lines {
		if len(line) == 0 {
			continue
		}
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "<key>") {
			err := parsePlaylistHeader(line, &playlist)
			if err != nil {
				return playlist, err
			}
		} else if strings.HasPrefix(line, "<array>") {
			playlist.TrackIDs = parseArray(lines[i+1 : len(lines)-1])
			return playlist, nil
		}
	}

	return playlist, nil
}

func main() {
	file, err := os.Open("Bibliothèque.xml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

}
