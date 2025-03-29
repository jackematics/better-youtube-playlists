export function getPlaylists() {
  const playlistListItems = localStorage.getItem("playlistListItems");
  if (!playlistListItems) {
    return [];
  } else {
    return JSON.parse(playlistListItems);
  }
}

export function addPlaylist(playlistListItem) {
  const playlistListItems = getPlaylists();

  localStorage.setItem(
    "playlistListItems",
    JSON.stringify([...playlistListItems, playlistListItem])
  );
}

export function removePlaylist(playlistId) {
  const playlistListItems = getPlaylists();
  playlistListItems.splice(
    playlistListItems.findIndex((item) => item.playlistId === playlistId),
    1
  );

  if (playlistListItems.length) {
    localStorage.setItem(
      "playlistListItems",
      JSON.stringify(playlistListItems)
    );
  } else {
    localStorage.clear();
  }
}
