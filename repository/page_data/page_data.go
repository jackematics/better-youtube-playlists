package page_data

import (
	"github.com/jackematics/better-youtube-playlists/model"
)

var IndexState = InitialiseState()

var NilPlaylist = model.Playlist{
	PlaylistId:    "default-playlist-id",
	PlaylistTitle: "No Playlist Selected",
	ChannelOwner:  "",
	TotalVideos:   0,
	Selected:      false,
	PlaylistItems: []model.PlaylistItem{},
}

func InitialiseState() model.Index {
	return model.Index{
		ModalState: model.Modal{
			Hidden:            true,
			ValidationMessage: "",
		},
		PlaylistListState: []model.Playlist{
			{
				PlaylistId:    "default-playlist-id",
				PlaylistTitle: "No Playlist Selected",
				ChannelOwner:  "",
				TotalVideos:   0,
				Selected:      false,
				PlaylistItems: []model.PlaylistItem{},
			},
		},
	}
}

func ToggleAddPlaylistModal() {
	IndexState.ModalState.Hidden = !IndexState.ModalState.Hidden
}

func AddPlaylist(playlist_model model.Playlist) {
	IndexState.PlaylistListState = append(IndexState.PlaylistListState, playlist_model)
}

func ResetAddPlaylistValidation() {
	IndexState.ModalState.ValidationMessage = ""
}

func SetValidationMessage(message string) {
	IndexState.ModalState.ValidationMessage = message
}

// returns true if the value is found and false otherwise
func FindPlaylist(playlist_id string) (*model.Playlist, bool) {
	playlist_list_state := IndexState.PlaylistListState
	for i := range playlist_list_state {
		if playlist_list_state[i].PlaylistId == playlist_id {
			return &playlist_list_state[i], true
		}
	}

	return nil, false
}

func SetSelectedPage(playlist_id string) bool {
	playlists := &IndexState.PlaylistListState

	selected_found := false
	for i := range *playlists {
		if (*playlists)[i].PlaylistId == playlist_id {
			(*playlists)[i].Selected = true
			selected_found = true
		} else {
			(*playlists)[i].Selected = false
		}
	}

	return selected_found
}

func GetPlaylistIndex(playlist_id string) int {
	playlists := IndexState.PlaylistListState

	for i := range playlists {
		if playlists[i].PlaylistId == playlist_id {
			return i
		}
	}

	return -1
}

func FindSelected() (*model.Playlist, bool) {
	playlists := &IndexState.PlaylistListState

	for i := range *playlists {
		playlist := &(*playlists)[i]
		if (*playlists)[i].Selected {
			return playlist, true
		}
	}

	return nil, false
}
