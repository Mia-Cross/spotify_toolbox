package itunes_library_parser

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsePlaylist(t *testing.T) {
	data := `
			<key>Name</key><string>♦ -Fall 2016- ♦</string>
			<key>Description</key><string></string>
			<key>Playlist ID</key><integer>34614</integer>
			<key>Playlist Persistent ID</key><string>C04A1AB49759EDD7</string>
			<key>All Items</key><true/>
			<key>Playlist Items</key>
			<array>
				<dict>
					<key>Track ID</key><integer>3036</integer>
				</dict>
				<dict>
					<key>Track ID</key><integer>6500</integer>
				</dict>
				<dict>
					<key>Track ID</key><integer>6502</integer>
				</dict>
				<dict>
					<key>Track ID</key><integer>9296</integer>
				</dict>
			</array>
`
	playlist, err := ParsePlaylist(strings.Split(data, "\n"))
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "♦ -Fall 2016- ♦", playlist.Name)
	assert.Equal(t, "", playlist.Description)
	assert.Equal(t, 34614, playlist.PlaylistID)
	assert.Equal(t, []int{3036, 6500, 6502, 9296}, playlist.TrackIDs)
}

//func TestParsePlaylists(t *testing.T) {
//	data := `
//	<key>Playlists</key>
//	<array>
//		<dict>
//			<key>Name</key><string>Bibliothèque</string>
//			<key>Description</key><string></string>
//			<key>Master</key><true/>
//			<key>Playlist ID</key><integer>14024</integer>
//			<key>Playlist Persistent ID</key><string>8D24B71203C8471A</string>
//			<key>Visible</key><false/>
//			<key>All Items</key><true/>
//			<key>Playlist Items</key>
//			<array>
//				<dict>
//					<key>Track ID</key><integer>3036</integer>
//				</dict>
//				<dict>
//					<key>Track ID</key><integer>6500</integer>
//				</dict>
//				<dict>
//					<key>Track ID</key><integer>6502</integer>
//				</dict>
//				<dict>
//					<key>Track ID</key><integer>9296</integer>
//				</dict>
//			</array>
//		</dict>
//		<dict>
//			<key>Name</key><string>Mémos vocaux</string>
//			<key>Description</key><string></string>
//			<key>Playlist ID</key><integer>35407</integer>
//			<key>Playlist Persistent ID</key><string>2E28BAEA6E4B2C88</string>
//			<key>Distinguished Kind</key><integer>17</integer>
//			<key>All Items</key><true/>
//		</dict>
//	</array>
//`
//}
