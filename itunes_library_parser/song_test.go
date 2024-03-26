package itunes_library_parser

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseSong(t *testing.T) {
	data := `
			<key>Track ID</key><integer>3036</integer>
			<key>Name</key><string>Washington Square</string>
			<key>Artist</key><string>Chinese Man</string>
			<key>Date Added</key><date>2016-04-25T21:52:15Z</date>
`
	song, err := ParseSong(strings.Split(data, "\n")[1:5])
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 3036, song.TrackID)
	assert.Equal(t, "Washington Square", song.Name)
	assert.Equal(t, "Chinese Man", song.Artist)
	if !song.DateAdded.Equal(time.Date(2016, 04, 25, 21, 52, 15, 0, time.UTC)) {
		t.Error("wrong date: expected \"2016-04-25 21:52:15 +0000 UTC\", got " + song.DateAdded.String())
	}
}

func TestParseSongs(t *testing.T) {
	data := `
	<dict>
		<key>3036</key>
		<dict>
			<key>Track ID</key><integer>3036</integer>
			<key>Name</key><string>Washington Square</string>
			<key>Artist</key><string>Chinese Man</string>
			<key>Date Added</key><date>2016-04-25T21:52:15Z</date>
		</dict>
		<key>3038</key>
		<dict>
			<key>Track ID</key><integer>3038</integer>
			<key>Name</key><string>Air War</string>
			<key>Artist</key><string>Crystal Castles</string>
			<key>Date Added</key><date>2017-04-30T21:32:16Z</date>
		</dict>
	</dict>
`
	songs, err := ParseSongs(strings.Split(data, "\n")[1:])
	if err != nil {
		t.Error(err)
	}

	song := songs[3036]
	assert.Equal(t, "Washington Square", song.Name)
	assert.Equal(t, "Chinese Man", song.Artist)
	if !song.DateAdded.Equal(time.Date(2016, 04, 25, 21, 52, 15, 0, time.UTC)) {
		t.Error("wrong date: expected \"2016-04-25 21:52:15 +0000 UTC\", got " + song.DateAdded.String())
	}

	song = songs[3038]
	assert.Equal(t, "Air War", song.Name)
	assert.Equal(t, "Crystal Castles", song.Artist)
	if !song.DateAdded.Equal(time.Date(2017, 04, 30, 21, 32, 16, 0, time.UTC)) {
		t.Error("wrong date: expected \"2017-04-30 21:32:16Z +0000 UTC\", got " + song.DateAdded.String())
	}
}
