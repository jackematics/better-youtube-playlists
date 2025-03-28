import { handleNextClick } from "./playlist-operations.js";

const videoContainerId = "video-container";
const videoContainerEl = document.getElementById(videoContainerId);

let player;

function onPlayerStateChange(event) {
  if (event.data === YT.PlayerState.ENDED) {
    handleNextClick();
  }
}

export function setPlayingVideo(videoId) {
  if (player) {
    player.loadVideoById(videoId);
  } else {
    player = new YT.Player(videoContainerId, {
      className: "w-full h-full",
      videoId,
      title: "youtube video",
      playerVars: {
        autoplay: 1,
        rel: 0,
        modestbranding: 1,
        origin: "https://www.youtube-nocookie.com",
      },
      events: {
        onStateChange: onPlayerStateChange,
      },
    });
  }
}

export function destroyPlayer() {
  if (player) {
    player.stopVideo();
    player.destroy();

    player = null;
  }

  videoContainerEl.innerHTML = "";
  const placeholderImgEl = document.createElement("img");
  placeholderImgEl.src = "/static/assets/logos/jackematica-logo.svg";
  placeholderImgEl.alt = "page-logo";
  placeholderImgEl.width = 210;
  placeholderImgEl.height = 210;
  videoContainerEl.appendChild(placeholderImgEl);
}
