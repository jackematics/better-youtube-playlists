package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jackematics/better-youtube-playlists/api"
	"github.com/jackematics/better-youtube-playlists/repository"
	"github.com/stretchr/testify/assert"
)

func TestModalHiddenByDefault(t *testing.T) {
	state := repository.GetPageState()

	assert.Equal(t, true, state.ModalState.Hidden)
}

func TestModalOpens(t *testing.T) {
	req, err := http.NewRequest("GET", "/toggle-add-playlist-modal", nil)
	state := repository.GetPageState()

	assert.Equal(t, nil, err)

	res_recorder := httptest.NewRecorder()
	http.HandlerFunc(api.ToggleAddPlaylistModalHandler).ServeHTTP(res_recorder, req)

	assert.Equal(t, false, state.ModalState.Hidden)
	assert.Equal(t, http.StatusOK, res_recorder.Code)
}
