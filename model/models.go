package model

type IndexModel struct {
	ModalState    ModalModel
	PlaylistState []PlaylistModel
}

type ModalModel struct {
	Hidden            bool
	ValidationMessage string
}

type PlaylistModel struct {
	PlaylistId    string
	PlaylistTitle string
	ChannelOwner  string
}
