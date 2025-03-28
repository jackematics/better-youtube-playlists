import { History } from "./history.js";
import { closeModal } from "./modal.js";
import { createPlaylistItem, highlightSelectedItem } from "./playlist-items.js";
import {
  resetOperationsState,
  setOriginalPlaylistItems,
} from "./playlist-operations.js";
import { destroyPlayer, setPlayingVideo } from "./youtube-embed.js";
import { addPlaylist, getPlaylists } from "./localStorage.js";

const playlistListItemsEl = document.getElementById("playlist-list-items");
const playlistTitleEl = document.getElementById("playlist-title");
const totalVideosEl = document.getElementById("total-videos");
const selectPlaylistValidationMessageEl = document.getElementById(
  "select-playlist-validation-message"
);

const playlistOperationsEl = document.getElementById("playlist-operations");

const loaderContainerEl = document.getElementById("loader-container");
const playlistItemsEl = document.getElementById("playlist-items");
const playlistIdInputEl = document.getElementById("playlist-id-input");
const modalValidationMessageEl = document.getElementById(
  "modal-validation-message"
);
const submitPlaylistBtnEl = document.getElementById("submit-playlist-button");

async function handlePlaylistClick(event, playlistId) {
  // Don't refetch when selecting same playlist
  if (playlistTitleEl.textContent === event.currentTarget.textContent) {
    return;
  }

  playlistItemsEl.innerHTML = "";
  playlistTitleEl.textContent = event.currentTarget.textContent;

  document.querySelectorAll(".playlist").forEach((playlist) => {
    playlist.classList.remove("bg-warm-orange");
  });
  event.currentTarget.classList.add("bg-warm-orange");

  // shows loading spinner
  loaderContainerEl.classList.remove("invisible");

  selectPlaylistValidationMessageEl.textContent = "";
  let validationMessage = "";

  try {
    History.clear();

    resetOperationsState();

    const response = await fetch(`/playlist-items/${playlistId}`);

    if ([500, 424].includes(response.status)) {
      validationMessage = "Internal server error";
      throw new Error(validationMessage);
    }

    if (response.status === 400) {
      validationMessage = "Invalid playlist ID";
      throw new Error(validationMessage);
    }

    const playlist = await response.json();

    totalVideosEl.textContent = `Video: 1 / ${playlist.totalVideos}`;

    if (playlist.items.length === 0) {
      validationMessage = "Playlist is empty";
      throw new Error(validationMessage);
    }

    playlistOperationsEl.classList.remove("invisible");

    playlistItemsEl.innerHTML = "";
    for (let i = 0; i < playlist.items.length; i++) {
      playlistItemsEl.appendChild(createPlaylistItem(i + 1, playlist.items[i]));
    }

    setOriginalPlaylistItems(Array.from(playlistItemsEl.children));
    highlightSelectedItem(document.getElementById(playlist.items[0].id));

    setPlayingVideo(playlist.items[0].id);
    History.add(playlist.items[0].id);
  } catch (err) {
    destroyPlayer();
    totalVideosEl.textContent = "";
    selectPlaylistValidationMessageEl.textContent = validationMessage;
    playlistOperationsEl.classList.add("invisible");
  }

  loaderContainerEl.classList.add("invisible");
}

function createPlaylistListItem(item) {
  const playlistListItem = document.createElement("li");
  playlistListItem.id = `${item.playlistId}`;
  playlistListItem.setAttribute("key", item.playlistId);
  playlistListItem.className = "mb-2 select-none";

  const option = document.createElement("option");
  option.className =
    "playlist mr-3 pl-2 pr-2 pb-1 pt-1 text-2xl text-left rounded-2xl cursor-pointer truncate font-medium text-white hover:bg-warm-orange-hover";
  option.title = item.playlistTitle;
  option.textContent = item.playlistTitle;

  option.addEventListener("click", (event) => {
    handlePlaylistClick(event, item.playlistId);
  });

  playlistListItem.appendChild(option);

  return playlistListItem;
}

async function handleAddPlaylist() {
  const playlistId = playlistIdInputEl.value;

  if (!playlistId) {
    modalValidationMessageEl.textContent = "Invalid playlist id";
    return;
  }

  try {
    const response = await fetch(`/playlist-metadata/${playlistId}`);

    if ([500, 424].includes(response.status)) {
      modalValidationMessageEl.textContent = "Internal server error";
      return;
    }

    if (response.status === 400) {
      modalValidationMessageEl.textContent = "Invalid playlist ID";
      return;
    }

    const newPlaylistListItem = await response.json();

    const playlistListItems = getPlaylists();

    if (
      playlistListItems
        .map((data) => data.playlistId)
        .includes(newPlaylistListItem.playlistId)
    ) {
      modalValidationMessageEl.textContent = "Duplicate playlists forbidden";
      return;
    }

    addPlaylist(newPlaylistListItem);

    closeModal();
  } catch (err) {
    console.log(err);
    if ([500, 424].includes(err.statusCode)) {
      modalValidationMessageEl.textContent = `Error fetching playlist data`;
      return;
    }

    if (err.statusCode === 400) {
      modalValidationMessageEl.textContent = `No playlist items returned for playlist`;
      return;
    }
  }

  renderList();
}

function renderList() {
  playlistListItemsEl.innerHTML = "";

  const playlists = getPlaylists();
  playlists.sort(
    (itemA, itemB) =>
      itemA.playlistTitle.charCodeAt(0) - itemB.playlistTitle.charCodeAt(0)
  );

  playlists.forEach((item) => {
    playlistListItemsEl.appendChild(createPlaylistListItem(item));
  });
}

document.addEventListener("DOMContentLoaded", renderList());

submitPlaylistBtnEl.addEventListener("click", handleAddPlaylist);
