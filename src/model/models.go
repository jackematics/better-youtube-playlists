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
	PlaylistId    string
	PlaylistTitle string
	ChannelOwner  string
	TotalVideos   int
	Selected      bool
	PlaylistItems []PlaylistItem
}
