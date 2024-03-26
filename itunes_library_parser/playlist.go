package itunes_library_parser

import (
	"strconv"
	"strings"
)

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

func ParsePlaylist(data []string) (Playlist, error) {
	playlist := Playlist{}

	for i, line := range data {
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "<key>") {
			err := parsePlaylistHeader(line, &playlist)
			if err != nil {
				return playlist, err
			}
		} else if strings.HasPrefix(line, "<array>") {
			playlist.TrackIDs = parseArray(data[i+1 : len(data)-1])
			return playlist, nil
		}
	}

	return playlist, nil
}
