package model

type Index struct {
	ModalState    Modal
	PlaylistState Playlist
}

type Modal struct {
	Hidden            bool
	ValidationMessage string
}

type Playlist struct {
	Playlists                 []PlaylistItem
	SelectedPlaylistItemIndex int
}

type PlaylistItem struct {
	PlaylistId    string
	PlaylistTitle string
	ChannelOwner  string
}
