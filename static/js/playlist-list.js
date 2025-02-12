function renderList() {
  const list = document.getElementById("playlist-list-items");
  list.innerHTML = "";
  const playlistListItems =
    JSON.parse(localStorage.getItem("playlistListItems")) || [];
  playlistListItems.forEach((item) => {
    const li = document.createElement("li");
    li.id = `li-${item.playlistId}`;
    li.setAttribute("key", item.playlistId);
    li.className = "mb-2";

    const option = document.createElement("option");
    option.className =
      "mr-3 pl-2 pr-2 pb-1 pt-1 text-2xl text-left rounded-2xl cursor-pointer truncate font-medium text-white hover:bg-warm-orange-hover";
    option.title = item.playlistTitle;
    option.textContent = item.playlistTitle;

    li.appendChild(option);
    list.appendChild(li);
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
      const response = await fetch(`/add-playlist/${playlistId}`);

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
