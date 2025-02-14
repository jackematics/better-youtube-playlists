package model

type Index struct {
	ModalState        Modal
	PlaylistListState []PlaylistMetadata
}

type Modal struct {
	Hidden            bool
	ValidationMessage string
}

type Item struct {
	Id           string `json:"id"`
	Title        string `json:"title"`
	ThumbnailUrl string `json:"thumbnailUrl"`
}

type Playlist struct {
	TotalVideos  int `json:"totalVideos"`
	Items        []Item `json:"items"`
}

type PlaylistMetadata struct {
	PlaylistId    string `json:"playlistId"`
	PlaylistTitle string `json:"playlistTitle"`
	ChannelOwner  string `json:"channelOwner"`
}
