package select_playlist

import (
	"html/template"
	"log"
	"net/http"

	"github.com/jackematics/better-youtube-playlists/repository/page_data"
)

func SetPlaylistDescriptionHandler(writer http.ResponseWriter, reader *http.Request) {
	playlist_id := reader.URL.Query().Get("playlist_id")

	selected_playlist_data, found := page_data.FindPlaylist(playlist_id)

	if !found {
		log.Println("No playlist description exists for id " + playlist_id)

		http.Error(writer, "No playlist description exists for "+playlist_id, http.StatusBadRequest)
	}

	log.Println("Selected description for playlist \"" + selected_playlist_data.PlaylistTitle + "\" from playlist_id \"" + playlist_id + "\"")
	tmpl := template.Must(template.ParseFiles("templates/playlist-description.html"))
	tmpl.ExecuteTemplate(writer, "playlist-description", selected_playlist_data)
}

func SetPlaylistItemsHandler(writer http.ResponseWriter, reader *http.Request) {
	playlist_id := reader.URL.Query().Get("playlist_id")

	selected_playlist_data, found := page_data.FindPlaylist(playlist_id)

	if !found {
		log.Println("No playlist items exist for id " + playlist_id)

		http.Error(writer, "No playlist items exist for "+playlist_id, http.StatusBadRequest)
	}

	log.Println("Selected items for playlist \"" + selected_playlist_data.PlaylistTitle + "\" from playlist_d \"" + playlist_id + "\"")
	tmpl := template.Must(template.ParseFiles("templates/playlist-items.html"))
	tmpl.ExecuteTemplate(writer, "playlist-items", selected_playlist_data)
}
