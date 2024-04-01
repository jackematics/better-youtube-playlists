package model

type Index struct {
	ModalState        Modal
	PlaylistListState []Playlist
}

type Modal struct {
	Hidden            bool
	ValidationMessage string
}

type Playlist struct {
	PlaylistId    string
	PlaylistTitle string
	ChannelOwner  string
	Selected      bool
}
