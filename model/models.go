package model

type Index struct {
	ModalState        Modal
	PlaylistListState []Playlist
}

type Modal struct {
	Hidden            bool
	ValidationMessage string
}

type PlaylistItem struct {
	Id           string
	Title        string
	ThumbnailUrl string
	Selected     bool
}

type Playlist struct {
	PlaylistId    string `json:"playlistId"`
	PlaylistTitle string `json:"playlistTitle"`
	ChannelOwner  string `json:"channelOwner"`
}
