import { closeModal } from "./modal.js";
import { highlightSelectedItem } from "./playlist-items.js";
import { setPlayingVideo } from "./youtube-embed.js";

async function handlePlaylistItemClick(event) {
  highlightSelectedItem(event.currentTarget);

  // scroll to centre of container
  const playlistItemsContainer = document.getElementById(
    "playlist-items-container"
  );
  playlistItemsContainer.scrollTo({
    top:
      event.currentTarget.offsetTop -
      (playlistItemsContainer.clientHeight / 2 - 51.2),
    behavior: "smooth",
  });

  setPlayingVideo(event.currentTarget.id);

  // set video currently playing in description
  const itemIndex = event.currentTarget.children[0].children[0].textContent;
  const totalVideos = document.getElementById("total-videos");
  const pSplit = totalVideos.textContent.split(" ");
  const totalVideoCount = pSplit[pSplit.length - 1];
  totalVideos.textContent = `Videos: ${itemIndex} / ${totalVideoCount}`;
}

async function createPlaylistItem(playlistItems, i) {
  const playlistItem = document.createElement("li");
  playlistItem.id = playlistItems[i].id;
  playlistItem.setAttribute("key", playlistItems[i].id);
  playlistItem.className =
    "playlist-items h-[3.2rem] pt-1 pb-1 pr-3 mr-2 ml-3 flex flex-row text-[1.75rem] text-left cursor-pointer text-white hover:bg-warm-orange-hover select-none";

  const indexContainer = document.createElement("div");
  indexContainer.className = "w-[4.5rem] flex justify-center";

  const index = document.createElement("p");
  index.textContent = i + 1;

  indexContainer.appendChild(index);

  const thumbnail = document.createElement("img");
  thumbnail.src = playlistItems[i].thumbnailUrl;
  thumbnail.alt = playlistItems[i].title;
  thumbnail.width = 70;
  thumbnail.height = 36;
  thumbnail.className = "ml-6 mr-2";

  const title = document.createElement("p");
  title.className = "pl-7 w-[67rem] truncate";
  title.textContent = playlistItems[i].title;

  playlistItem.appendChild(indexContainer);
  playlistItem.appendChild(thumbnail);
  playlistItem.appendChild(title);

  playlistItem.addEventListener("click", handlePlaylistItemClick);

  return playlistItem;
}

async function handlePlaylistClick(event, playlistId) {
  const playlistTitle = document.getElementById("playlist-title");

  // Don't refetch when selecting same playlist
  if (playlistTitle.textContent === event.currentTarget.textContent) {
    return;
  }

  // clears playlist items
  const playlistItems = document.getElementById("playlist-items");
  playlistItems.innerHTML = "";

  // sets new title
  playlistTitle.textContent = event.currentTarget.textContent;

  // highlights playlist
  document.querySelectorAll(".playlist").forEach((playlist) => {
    playlist.classList.remove("bg-warm-orange");
  });
  event.currentTarget.classList.add("bg-warm-orange");

  // shows loading spinner
  const loader = document.getElementById("loader-container");
  loader.classList.remove("invisible");

  try {
    const response = await fetch(`/playlist-items/${playlistId}`);

    if ([500, 424].includes(response.status)) {
      validationMessage.textContent = "Internal server error";
      return;
    }

    if (response.status === 400) {
      validationMessage.textContent = "Invalid playlist ID";
      return;
    }

    const playlist = await response.json();

    // populate playlist description

    document.getElementById(
      "total-videos"
    ).textContent = `Videos: 1 / ${playlist.totalVideos}`;

    const channelOwner = JSON.parse(
      localStorage.getItem("playlistListItems")
    ).find((storedItem) => storedItem.playlistId === playlistId).channelOwner;

    if (channelOwner) {
      document.getElementById("channel-owner").textContent = channelOwner;
    }

    // show operations

    document
      .getElementById("playlist-operations")
      .classList.remove("invisible");

    // create list items
    for (let i = 0; i < playlist.items.length; i++) {
      playlistItems.appendChild(await createPlaylistItem(playlist.items, i));
    }

    highlightSelectedItem(document.getElementById(playlist.items[0].id));
    // play first item in list
    setPlayingVideo(playlist.items[0].id);
  } catch (err) {
    console.log("cheese", err);

    // TODO: handle error getting playlist items? toast?
  }

  loader.classList.add("invisible");
}

function createPlaylistListItem(item) {
  const playlistListItem = document.createElement("li");
  playlistListItem.id = `li-${item.playlistId}`;
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
  const playlistId = document.getElementById("playlist-id-input").value;
  const validationMessage = document.getElementById("validation-message");

  if (!playlistId) {
    validationMessage.textContent = "Invalid playlist id";
    return;
  }

  try {
    const response = await fetch(`/playlist-metadata/${playlistId}`);

    if ([500, 424].includes(response.status)) {
      validationMessage.textContent = "Internal server error";
      return;
    }

    if (response.status === 400) {
      validationMessage.textContent = "Invalid playlist ID";
      return;
    }

    const newPlaylistListItem = await response.json();

    const playlistListItems =
      JSON.parse(localStorage.getItem("playlistListItems")) || [];

    if (
      playlistListItems
        .map((data) => data.playlistId)
        .includes(newPlaylistListItem.playlistId)
    ) {
      validationMessage.textContent = "Duplicate playlists forbidden";
      return;
    }

    localStorage.setItem(
      "playlistListItems",
      JSON.stringify([...playlistListItems, newPlaylistListItem])
    );

    closeModal();
  } catch (err) {
    console.log(err);
    if ([500, 424].includes(err.statusCode)) {
      validationMessage.textContent = `Error fetching playlist data`;
      return;
    }

    if (err.statusCode === 400) {
      validationMessage.textContent = `No playlist items returned for playlist`;
      return;
    }
  }

  renderList();
}

function renderList() {
  const playlistList = document.getElementById("playlist-list-items");
  playlistList.innerHTML = "";

  const playlistListItems =
    JSON.parse(localStorage.getItem("playlistListItems")) || [];

  playlistListItems.forEach((item) => {
    playlistList.appendChild(createPlaylistListItem(item));
  });
}

document.addEventListener("DOMContentLoaded", renderList());

document
  .getElementById("submit-playlist-button")
  .addEventListener("click", handleAddPlaylist);
