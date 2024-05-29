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
