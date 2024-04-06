package model

type Index struct {
	ModalState        Modal
	PlaylistListState []Playlist
}

type Modal struct {
	Hidden            bool
	ValidationMessage string
}

type Thumbnail struct {
	Url    string
	Width  int
	Height int
}

type PlaylistItem struct {
	Id        string
	Title     string
	Thumbnail Thumbnail
}

type Playlist struct {
	PlaylistId    string
	PlaylistTitle string
	ChannelOwner  string
	TotalVideos   int
	Selected      bool
	PlaylistItems []PlaylistItem
}
