import { History } from "./history.js";
import { removePlaylist } from "./localStorage.js";

let RANDOMISE = false;
let SHUFFLE = false;
let LOOP = false;
let ORIGINAL_PLAYLIST_ITEMS = [];

const playlistEl = document.getElementById("playlist-list-items");
const playlistItemsEl = document.getElementById("playlist-items");
const totalVideosEl = document.getElementById("total-videos");

const previousEl = document.getElementById("previous");
const nextEl = document.getElementById("next");
const randomiseEl = document.getElementById("randomise");
const shuffleEl = document.getElementById("shuffle");
const loopEl = document.getElementById("loop");
const removePlaylistEl = document.getElementById("remove-playlist");

function handlePreviousClick() {
  const historyPrev = History.getPreviousVideoId();

  let prevVideo = historyPrev
    ? document.getElementById(historyPrev)
    : playlistItemsEl.querySelector(".bg-warm-orange").previousElementSibling;

  // loop to end if configured
  if (!prevVideo && LOOP) {
    prevVideo = playlistItemsEl.lastElementChild;
  }

  if (prevVideo) {
    prevVideo.click();
  }
}

export function handleNextClick() {
  const currentVideo = playlistItemsEl.querySelector(".bg-warm-orange");

  let nextVideo;
  if (RANDOMISE) {
    const currentItemIndex = Number.parseInt(
      currentVideo.children[0].children[0].textContent
    );
    const totalVideosSplit = totalVideosEl.textContent.split(" ");
    const totalVideoCount = Number.parseInt(
      totalVideosSplit[totalVideosSplit.length - 1]
    );

    let nextLiIndex;
    do {
      nextLiIndex = Math.floor(Math.random() * totalVideoCount);
    } while (nextLiIndex === currentItemIndex);

    nextVideo = playlistItemsEl.children[nextLiIndex];
  } else {
    nextVideo = currentVideo.nextElementSibling;

    // loop to beginning if configured
    if (!nextVideo && LOOP) {
      nextVideo = playlistItemsEl.children[0];
    }
  }

  if (nextVideo) {
    nextVideo.click();
  }
}

function handleRandomise() {
  RANDOMISE = !RANDOMISE;

  if (RANDOMISE) {
    randomiseEl.classList.remove("bg-white");
    randomiseEl.classList.add("bg-orange-highlight");
  } else {
    History.clear();

    randomiseEl.classList.remove("bg-orange-highlight");
    randomiseEl.classList.add("bg-white");
  }
}

export function setOriginalPlaylistItems(playlistItems) {
  ORIGINAL_PLAYLIST_ITEMS = playlistItems;
}

function handleShuffle() {
  SHUFFLE = !SHUFFLE;

  if (SHUFFLE) {
    shuffleEl.classList.remove("bg-white");
    shuffleEl.classList.add("bg-orange-highlight");

    const newPlaylistItemIds = [];
    const oldPlaylistItems = Array.from(playlistItemsEl.children);

    // shuffle playlist
    while (oldPlaylistItems.length) {
      const randomIndex = Math.floor(Math.random() * oldPlaylistItems.length);
      newPlaylistItemIds.push(oldPlaylistItems.splice(randomIndex, 1)[0]);
    }

    playlistItemsEl.innerHTML = "";

    for (let i = 0; i < newPlaylistItemIds.length; i++) {
      const indexEl = newPlaylistItemIds[i].querySelector("p");
      indexEl.textContent = i + 1;

      playlistItemsEl.appendChild(newPlaylistItemIds[i]);
    }

    playlistItemsEl.children[0].click();
  } else {
    History.clear();

    // shuffle button styling
    shuffleEl.classList.remove("bg-orange-highlight");
    shuffleEl.classList.add("bg-white");

    // unshuffle playlist
    playlistItemsEl.innerHTML = "";
    for (let i = 0; i < ORIGINAL_PLAYLIST_ITEMS.length; i++) {
      const indexEl = ORIGINAL_PLAYLIST_ITEMS[i].querySelector("p");
      indexEl.textContent = i + 1;

      playlistItemsEl.appendChild(ORIGINAL_PLAYLIST_ITEMS[i]);
    }

    playlistItemsEl.children[0].click();
  }
}

function handleLoop() {
  LOOP = !LOOP;

  if (LOOP) {
    loopEl.classList.remove("bg-white");
    loopEl.classList.add("bg-orange-highlight");
  } else {
    loopEl.classList.remove("bg-orange-highlight");
    loopEl.classList.add("bg-white");
  }
}

function handleRemovePlaylist() {
  const selectedPlaylistId = [...playlistEl.querySelectorAll("li")].find((li) =>
    li.querySelector("option.bg-warm-orange")
  ).id;

  removePlaylist(selectedPlaylistId);

  location.reload();
}

export function resetOperationsState() {
  if (RANDOMISE) {
    handleRandomise();
  }

  if (SHUFFLE) {
    handleShuffle();
  }

  if (LOOP) {
    handleLoop();
  }
}

previousEl.addEventListener("click", handlePreviousClick);
nextEl.addEventListener("click", handleNextClick);
randomiseEl.addEventListener("click", handleRandomise);
shuffleEl.addEventListener("click", handleShuffle);
loopEl.addEventListener("click", handleLoop);
removePlaylistEl.addEventListener("click", handleRemovePlaylist);
