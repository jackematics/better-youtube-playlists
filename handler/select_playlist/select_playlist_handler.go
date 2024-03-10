package select_playlist

import (
	"html/template"
	"log"
	"net/http"

	"github.com/jackematics/better-youtube-playlists/model"
	"github.com/jackematics/better-youtube-playlists/repository/page_data_repository"
)

func SetDescriptionHandler(writer http.ResponseWriter, reader *http.Request) {
	playlist_id := reader.URL.Query().Get("playlist_id")

	var selected_playlist_data model.PlaylistModel

	playlist_state := page_data_repository.IndexState.PlaylistState
	for i := range playlist_state {
		if playlist_state[i].PlaylistId == playlist_id {
			selected_playlist_data = playlist_state[i]
		}
	}

	if selected_playlist_data == (model.PlaylistModel{}) {
		log.Println("No playlist exists in state with id " + playlist_id)
		log.Printf("Current playlist state: %v\n", playlist_state)

		http.Error(writer, "Invalid playlist state: "+playlist_id, http.StatusBadRequest)
	}

	log.Println("Selected playlist \"" + selected_playlist_data.PlaylistTitle + "\" from playlist_id \"" + playlist_id + "\"")
	tmpl := template.Must(template.ParseFiles("templates/playlist-description.html"))
	tmpl.ExecuteTemplate(writer, "playlist-description", selected_playlist_data)
}
