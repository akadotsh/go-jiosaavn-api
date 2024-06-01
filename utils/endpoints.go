package utils

type SearchT struct {
	All       string
	Songs     string
	Albums    string
	Artists   string
	Playlists string
}

type SongsT struct {
	ID          string
	Link        string
	Suggestions string
	Lyrics      string
	Station     string
}

type AlbumsT struct {
	ID   string
	Link string
}


type ArtistsT struct{
	ID string
	Link string
	Songs string
	Albums string
}

var Search = SearchT{
	All:       "autocomplete.get",
	Songs:     "search.getResults",
	Albums:    "search.getAlbumResults",
	Artists:   "search.getArtistResults",
	Playlists: "search.getPlaylistResults",
}

var Songs = SongsT{
	ID:          "song.getDetails",
	Link:        "webapi.get",
	Suggestions: "webradio.getSong",
	Lyrics:      "lyrics.getLyrics",
	Station:     "webradio.createEntityStation",
}


type PlayListT struct{
	ID string
	Link string
}


var Album = AlbumsT{
	ID:   "content.getAlbumDetails",
	Link: "webapi.get",
}

var Artists =ArtistsT{
	ID: "artist.getArtistPageDetails",
	Link: "webapi.get",
	Songs: "artist.getArtistMoreSong",
	Albums: "artist.getArtistMoreAlbum",
}

var Playlist = PlayListT{
	ID: "playlist.getDetails",
	Link: "webapi.get", 
}