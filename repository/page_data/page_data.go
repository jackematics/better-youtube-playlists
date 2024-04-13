package page_data

import (
	"errors"

	"github.com/jackematics/better-youtube-playlists/model"
)

var IndexState = InitialiseState()

func InitialiseState() model.Index {
	return model.Index{
		ModalState: model.Modal{
			Hidden:            true,
			ValidationMessage: "",
		},
		PlaylistListState: []model.Playlist{},
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

func SetSelectedPlaylist(playlist_id string) bool {
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

func GetSelectedPlaylistIndex() int {
	selected_playlist_index := -1
	for i, playlist := range IndexState.PlaylistListState {
		if playlist.Selected {
			selected_playlist_index = i
			break
		}
	}

	return selected_playlist_index
}

func SetSelectedPlaylistItem(playlist_item_id string, selected_playlist_index int) int {
	if selected_playlist_index == -1 {
		return -1
	}

	selected_playlist_item_index := -1
	playlist_items_ref := &IndexState.PlaylistListState[selected_playlist_index].PlaylistItems
	for i := range *playlist_items_ref {
		if (*playlist_items_ref)[i].Id == playlist_item_id {
			(*playlist_items_ref)[i].Selected = true
			selected_playlist_item_index = i
		} else {
			(*playlist_items_ref)[i].Selected = false
		}
	}

	return selected_playlist_item_index
}

func GetSelectedPlaylistItem(playlist_item_id string) (*model.PlaylistItem, error) {
	selected_playlist_index := GetSelectedPlaylistIndex()

	if selected_playlist_index == -1 {
		return nil, errors.New("no playlist selected")
	}

	playlist_items_ref := &IndexState.PlaylistListState[selected_playlist_index].PlaylistItems
	for i := range *playlist_items_ref {
		if (*playlist_items_ref)[i].Id == playlist_item_id {
			return &(*playlist_items_ref)[i], nil
		}
	}

	return nil, errors.New("no playlist item found with id: " + playlist_item_id)
}
