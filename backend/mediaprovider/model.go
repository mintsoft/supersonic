package mediaprovider

type Album struct {
	ID          string
	CoverArtID  string
	Name        string
	Duration    int
	ArtistIDs   []string
	ArtistNames []string
	Year        int
	Genres      []string
	TrackCount  int
	Favorite    bool
}

type AlbumWithTracks struct {
	Album
	Tracks []*Track
}

type AlbumInfo struct {
	Notes         string
	LastFmUrl     string
	MusicBrainzID string
}

type Artist struct {
	ID         string
	CoverArtID string
	Name       string
	Favorite   bool
	AlbumCount int
}

type ArtistWithAlbums struct {
	Artist
	Albums []*Album
}

type ArtistInfo struct {
	Biography      string
	LastFMUrl      string
	ImageURL       string
	SimilarArtists []*Artist
}

type Genre struct {
	Name       string
	AlbumCount int
	TrackCount int
}

type Track struct {
	ID          string
	CoverArtID  string
	ParentID    string
	Name        string
	Duration    int
	TrackNumber int
	DiscNumber  int
	Genre       string
	ArtistIDs   []string
	ArtistNames []string
	Album       string
	AlbumID     string
	Year        int
	Rating      int
	Favorite    bool
	Size        int64
	PlayCount   int
	FilePath    string
	BitRate     int
}

type Playlist struct {
	ID          string
	CoverArtID  string
	Name        string
	Description string
	Public      bool
	Owner       string
	Duration    int
	TrackCount  int
}

type PlaylistWithTracks struct {
	Playlist
	Tracks []*Track
}

type ContentType int

const (
	ContentTypeAlbum ContentType = iota
	ContentTypeArtist
	ContentTypeTrack
	ContentTypePlaylist
	ContentTypeGenre
)

func (c ContentType) String() string {
	switch c {
	case ContentTypeAlbum:
		return "Album"
	case ContentTypeArtist:
		return "Artist"
	case ContentTypeTrack:
		return "Track"
	case ContentTypePlaylist:
		return "Playlist"
	case ContentTypeGenre:
		return "Genre"
	default:
		return "Unknown"
	}
}

type SearchResult struct {
	Name    string
	ID      string
	CoverID string
	Type    ContentType

	// for Album / Playlist: track count
	//     Artist / Genre: album count
	//     Track: length (seconds)
	Size int

	// Unset for ContentTypes Artist, Playlist, and Genre
	ArtistName string
}
