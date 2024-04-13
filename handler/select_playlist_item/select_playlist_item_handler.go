package select_playlist_item

import (
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/jackematics/better-youtube-playlists/helper/func_map"
	"github.com/jackematics/better-youtube-playlists/repository/page_data"
)

func HighlightSelectedPlaylistItemHandler(writer http.ResponseWriter, reader *http.Request) {
	url_parts := strings.Split(reader.URL.Path, "/")
	playlist_item_id := url_parts[len(url_parts)-1]

	selected_playlist_index := page_data.GetSelectedPlaylistIndex()
	if selected_playlist_index == -1 {
		http.Error(writer, "No playlist selected", http.StatusBadRequest)
		return
	}

	selected_item_index := page_data.SetSelectedPlaylistItem(playlist_item_id, selected_playlist_index)
	if selected_item_index == -1 {
		http.Error(writer, "No playlist item with id "+playlist_item_id, http.StatusBadRequest)
		return
	}

	tmpl := template.Must(template.New("playlist-items").Funcs(func_map.PageFuncs).ParseFiles("templates/playlist-items.html", "templates/playlist-item.html"))
	tmpl.ExecuteTemplate(writer, "playlist-items", page_data.IndexState.PlaylistListState[selected_playlist_index])
}

func PlaySelectedPlaylistItem(writer http.ResponseWriter, reader *http.Request) {
	url_parts := strings.Split(reader.URL.Path, "/")
	playlist_item_id := url_parts[len(url_parts)-1]

	selected_playlist_item, err := page_data.GetSelectedPlaylistItem(playlist_item_id)

	if err != nil {
		log.Println("error retrieving selected playlist item: ", err)
		http.Error(writer, "No playlist item with id "+playlist_item_id, http.StatusBadRequest)
		return
	}

	tmpl := template.Must(template.New("youtube-embed").Funcs(func_map.PageFuncs).ParseFiles("templates/youtube-embed.html"))
	tmpl.ExecuteTemplate(writer, "youtube-embed", selected_playlist_item)
}
