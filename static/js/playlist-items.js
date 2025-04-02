import { History } from "./history.js";
import { setPlayingVideo } from "./youtube-embed.js";
import { CurrentVideoState } from "./playlist-description.js";

const playlistItemsContainerEl = document.getElementById(
  "playlist-items-container"
);

export function highlightSelectedItem(selected) {
  document.querySelectorAll(".playlist-items").forEach((item) => {
    item.classList.remove("bg-warm-orange");
  });

  selected.classList.add("bg-warm-orange");
}

function handlePlaylistItemClick(event) {
  highlightSelectedItem(event.currentTarget);

  // scroll to centre of container
  playlistItemsContainerEl.scrollTo({
    top:
      event.currentTarget.offsetTop -
      (playlistItemsContainerEl.clientHeight / 2 - 51.2),
    behavior: "smooth",
  });

  setPlayingVideo(event.currentTarget.id);

  // set video currently playing in description
  CurrentVideoState.setState({
    currentIndex: Number.parseInt(
      event.currentTarget.children[0].children[0].textContent
    ),
  });
  CurrentVideoState.render();

  // adds video to history of played videos
  History.add(event.currentTarget.id);
}

export function createPlaylistItem(i, playlistItem) {
  const playlistItemEl = document.createElement("li");
  playlistItemEl.id = playlistItem.id;
  playlistItemEl.setAttribute("key", playlistItem.id);
  playlistItemEl.className =
    "playlist-items h-[3.2rem] pt-1 pb-1 pr-3 mr-2 ml-3 flex flex-row text-[1.75rem] text-left cursor-pointer text-white hover:bg-warm-orange-hover select-none";

  const indexContainer = document.createElement("div");
  indexContainer.className = "w-[4.5rem] flex justify-center";

  const index = document.createElement("p");
  index.textContent = i;

  indexContainer.appendChild(index);

  const thumbnail = document.createElement("img");
  thumbnail.src = playlistItem.thumbnailUrl;
  thumbnail.alt = playlistItem.title;
  thumbnail.width = 70;
  thumbnail.height = 36;
  thumbnail.className = "ml-6 mr-2";

  const title = document.createElement("p");
  title.className = "pl-7 w-[67rem] truncate";
  title.textContent = playlistItem.title;

  playlistItemEl.appendChild(indexContainer);
  playlistItemEl.appendChild(thumbnail);
  playlistItemEl.appendChild(title);

  playlistItemEl.addEventListener("click", handlePlaylistItemClick);

  return playlistItemEl;
}
