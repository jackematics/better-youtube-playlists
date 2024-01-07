package test

import (
	"testing"

	"github.com/jackematics/better-youtube-playlists/repository"
	"github.com/stretchr/testify/assert"
)

func TestModalHiddenByDefault(t *testing.T) {
	state := repository.GetPageState()

	assert.Equal(t, true, state.ModalState.Hidden)
}
