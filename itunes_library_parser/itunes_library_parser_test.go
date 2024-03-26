package itunes_library_parser

//func TestParseLibrary(t *testing.T) {
//	data := `
//<dict>
//	<key>Major Version</key><integer>1</integer>
//	<key>Minor Version</key><integer>1</integer>
//	<key>Date</key><date>2022-10-22T10:13:41Z</date>
//	<key>Application Version</key><string>12.12.2.2</string>
//	<key>Features</key><integer>5</integer>
//	<key>Show Content Ratings</key><true/>
//	<key>Music Folder</key><string>file://localhost/C:/Users/pepon/Music/iTunes/iTunes%20Media/</string>
//	<key>Library Persistent ID</key><string>559DC2EE171963AC</string>
//	<key>Tracks</key>
//	<dict>
//		<key>3036</key>
//		<dict>
//			<key>Track ID</key><integer>3036</integer>
//			<key>Name</key><string>Washington Square</string>
//			<key>Artist</key><string>Chinese Man</string>
//			<key>Album Artist</key><string>Chinese Man</string>
//			<key>Album</key><string>The Groove Sessions</string>
//			<key>Kind</key><string>Fichier audio MPEG</string>
//			<key>Size</key><integer>8064503</integer>
//			<key>Total Time</key><integer>251663</integer>
//			<key>Date Modified</key><date>2016-02-18T00:35:35Z</date>
//			<key>Date Added</key><date>2016-04-25T21:52:15Z</date>
//			<key>Bit Rate</key><integer>256</integer>
//			<key>Sample Rate</key><integer>44100</integer>
//			<key>Play Count</key><integer>9</integer>
//			<key>Play Date</key><integer>3714796367</integer>
//			<key>Play Date UTC</key><date>2021-09-18T05:52:47Z</date>
//			<key>Normalization</key><integer>935</integer>
//			<key>Sort Album</key><string>Groove Sessions</string>
//			<key>Sort Album Artist</key><string>Chinese Man</string>
//			<key>Persistent ID</key><string>B26E87A0D6E21820</string>
//			<key>Track Type</key><string>File</string>
//			<key>Location</key><string>file://localhost/C:/Users/pepon/Music/Chinese%20Man/Washington%20Square.mp3</string>
//			<key>File Folder Count</key><integer>-1</integer>
//			<key>Library Folder Count</key><integer>-1</integer>
//		</dict>
//		<key>3038</key>
//		<dict>
//			<key>Track ID</key><integer>3038</integer>
//			<key>Name</key><string>Air War</string>
//			<key>Artist</key><string>Crystal Castles</string>
//			<key>Album Artist</key><string>Crystal Castles</string>
//			<key>Album</key><string>Crystal Castles</string>
//			<key>Kind</key><string>Fichier audio MPEG</string>
//			<key>Size</key><integer>4090460</integer>
//			<key>Total Time</key><integer>255373</integer>
//			<key>Year</key><integer>2008</integer>
//			<key>Date Modified</key><date>2014-04-04T22:50:06Z</date>
//			<key>Date Added</key><date>2016-04-25T21:52:15Z</date>
//			<key>Bit Rate</key><integer>128</integer>
//			<key>Sample Rate</key><integer>44100</integer>
//			<key>Play Count</key><integer>1</integer>
//			<key>Play Date</key><integer>3552684423</integer>
//			<key>Play Date UTC</key><date>2016-07-29T22:47:03Z</date>
//			<key>Persistent ID</key><string>C512FD40638B07D8</string>
//			<key>Track Type</key><string>File</string>
//			<key>Location</key><string>file://localhost/C:/Users/pepon/Music/Crystal%20Castles/Air%20War.mp3</string>
//			<key>File Folder Count</key><integer>-1</integer>
//			<key>Library Folder Count</key><integer>-1</integer>
//		</dict>
//	</dict>
//</dict>
//`
//
//}
