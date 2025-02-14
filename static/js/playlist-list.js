import { closeModal } from "./modal.js";

function renderList() {
  const playlistList = document.getElementById("playlist-list-items");
  playlistList.innerHTML = "";

  const playlistListItems =
    JSON.parse(localStorage.getItem("playlistListItems")) || [];
  playlistListItems.forEach((item) => {
    const playlistListItem = document.createElement("li");
    playlistListItem.id = `li-${item.playlistId}`;
    playlistListItem.setAttribute("key", item.playlistId);
    playlistListItem.className = "mb-2";

    const option = document.createElement("option");
    option.className =
      "playlist mr-3 pl-2 pr-2 pb-1 pt-1 text-2xl text-left rounded-2xl cursor-pointer truncate font-medium text-white hover:bg-warm-orange-hover";
    option.title = item.playlistTitle;
    option.textContent = item.playlistTitle;

    option.addEventListener("click", async function () {
      const playlistTitle = document.getElementById("playlist-title");
      playlistTitle.textContent = item.playlistTitle;

      document.querySelectorAll(".playlist").forEach((playlist) => {
        playlist.classList.remove("bg-warm-orange");
      });
      option.classList.add("bg-warm-orange");

      try {
        const response = await fetch(`/playlist-items/${item.playlistId}`);

        if ([500, 424].includes(response.status)) {
          validationMessage.textContent = "Internal server error";
          return;
        }

        if (response.status === 400) {
          validationMessage.textContent = "Invalid playlist ID";
          return;
        }

        const newPlaylist = await response.json();

        document.getElementById(
          "total-videos"
        ).textContent = `Videos: ${newPlaylist.totalVideos}`;

        const channelOwner = JSON.parse(
          localStorage.getItem("playlistListItems")
        ).find(
          (storedItem) => storedItem.playlistId === item.playlistId
        ).channelOwner;

        if (channelOwner) {
          document.getElementById("channel-owner").textContent = channelOwner;
        }

        const playlistItems = document.getElementById("playlist-items");
        playlistItems.innerHTML = "";

        // create list item
        for (let i = 0; i < newPlaylist.items.length; i++) {
          const playlistItem = document.createElement("li");
          playlistItem.id = newPlaylist.items[i].id;
          playlistItem.setAttribute("key", newPlaylist.items[i].id);
          playlistItem.className =
            "h-[3.2rem] pt-1 pb-1 pr-3 mr-2 ml-3 flex flex-row text-[1.75rem] text-left cursor-pointer text-white hover:bg-warm-orange-hover";

          const indexContainer = document.createElement("div");
          indexContainer.className = "w-[4.5rem] flex justify-center";

          const index = document.createElement("p");
          index.textContent = i + 1;

          indexContainer.appendChild(index);

          const thumbnail = document.createElement("img");
          thumbnail.src = newPlaylist.items[i].thumbnailUrl;
          thumbnail.alt = newPlaylist.items[i].title;
          thumbnail.width = 70;
          thumbnail.height = 36;
          thumbnail.className = "ml-6 mr-2";

          const title = document.createElement("p");
          title.className = "pl-7 w-[67rem] truncate";
          title.textContent = newPlaylist.items[i].title;

          playlistItem.appendChild(indexContainer);
          playlistItem.appendChild(thumbnail);
          playlistItem.appendChild(title);

          playlistItems.appendChild(playlistItem);
        }
      } catch (err) {
        console.log("cheese", err);

        // handle error getting playlist items toast?
      }
    });

    playlistListItem.appendChild(option);
    playlistList.appendChild(playlistListItem);
  });
}

document.addEventListener("DOMContentLoaded", renderList());

document
  .getElementById("submit-playlist-button")
  .addEventListener("click", async function () {
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
  });
