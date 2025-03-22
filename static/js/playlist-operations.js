import { History } from "./history.js";

let RANDOMISE = false;

const playlistItemsEl = document.getElementById("playlist-items");
const totalVideosEl = document.getElementById("total-videos");

const previousEl = document.getElementById("previous");
const nextEl = document.getElementById("next");
const randomiseEl = document.getElementById("randomise");

function handlePreviousClick() {
  const historyPrev = History.getPreviousVideoId();

  const prevVideo = historyPrev
    ? document.getElementById(historyPrev)
    : playlistItemsEl.querySelector(".bg-warm-orange").previousElementSibling;

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
  }

  if (nextVideo) {
    nextVideo.click();
  }
}

function handleRandomise(event) {
  RANDOMISE = !RANDOMISE;

  if (RANDOMISE) {
    event.currentTarget.classList.remove("bg-white");
    event.currentTarget.classList.add("bg-orange-highlight");
  } else {
    History.clear();

    event.currentTarget.classList.remove("bg-orange-highlight");
    event.currentTarget.classList.add("bg-white");
  }
}

previousEl.addEventListener("click", handlePreviousClick);
nextEl.addEventListener("click", handleNextClick);
randomiseEl.addEventListener("click", handleRandomise);
