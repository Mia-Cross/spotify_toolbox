package itunes_library_parser

import (
	"errors"
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

func ParseSong(data []string) (Song, error) {
	song := Song{}
	var err error

	for _, line := range data {
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

func ParseSongs(data []string) (map[int]Song, error) {
	var songs map[int]Song

	for i, line := range data {
		// tmp fix
		if line == "\t<dict>" {
			continue
		}

		// extract song ID
		if !strings.Contains(line, "<key>") {
			return nil, errors.New("expected a key here")
		}
		id, _ := readKey(strings.TrimSpace(line))
		songID, err := strconv.Atoi(id)
		if err != nil {
			return nil, err
		}
		i++

		if strings.TrimSpace(data[i]) != "<dict>" {
			return nil, errors.New("expected the beginning of a dict here")
		}
		end, err := getDictLimit(data[i:])
		if err != nil {
			return nil, err
		}

		song, err := ParseSong(data[i+1 : end])
		if err != nil {
			return nil, err
		}

		songs[songID] = song
		i = end
	}

	return songs, nil
}
