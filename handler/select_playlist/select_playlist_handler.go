package select_playlist

import (
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/jackematics/better-youtube-playlists/helper/func_map"
	"github.com/jackematics/better-youtube-playlists/repository/page_data"
)

func SetPlaylistDescriptionHandler(writer http.ResponseWriter, reader *http.Request) {
	playlist_id := reader.URL.Query().Get("playlist_id")

	selected_playlist_data, found := page_data.FindPlaylist(playlist_id)

	if !found {
		log.Println("No playlist description exists for id " + playlist_id)

		http.Error(writer, "No playlist description exists for "+playlist_id, http.StatusBadRequest)
		return
	}

	log.Println("Selected description for playlist \"" + selected_playlist_data.PlaylistTitle + "\" from playlist_id \"" + playlist_id + "\"")
	tmpl := template.Must(template.ParseFiles("templates/playlist-description.html"))
	tmpl.ExecuteTemplate(writer, "playlist-description", selected_playlist_data)
}

func HighlightSelectedPlaylist(writer http.ResponseWriter, reader *http.Request) {
	url_parts := strings.Split(reader.URL.Path, "/")
	playlist_id := url_parts[len(url_parts)-1]

	ok := page_data.SetSelectedPage(playlist_id)

	if !ok {
		http.Error(writer, "No playlists exist for "+playlist_id, http.StatusBadRequest)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/playlist-list-item.html", "templates/playlist-list.html"))
	tmpl.ExecuteTemplate(writer, "playlist-list", page_data.IndexState.PlaylistListState)
}

func PopulatePlaylistItems(writer http.ResponseWriter, reader *http.Request) {
	url_parts := strings.Split(reader.URL.Path, "/")
	playlist_id := url_parts[len(url_parts)-1]

	selectedPlaylistIndex := page_data.GetPlaylistIndex(playlist_id)

	if selectedPlaylistIndex == -1 {
		http.Error(writer, "No playlist selected", http.StatusBadRequest)
		return
	}

	tmpl := template.Must(template.New("playlist-items").Funcs(func_map.Index).ParseFiles("templates/playlist-items.html", "templates/playlist-item.html"))
	tmpl.ExecuteTemplate(writer, "playlist-items", page_data.IndexState.PlaylistListState[selectedPlaylistIndex])
}
