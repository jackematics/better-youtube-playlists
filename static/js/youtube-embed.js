import { highlightSelectedItem } from "./playlist-items.js";

let player;

function onPlayerStateChange(event) {
  if (event.data === YT.PlayerState.ENDED) {
    const nextVideo = document
      .getElementById("playlist-items")
      .querySelector(".bg-warm-orange").nextElementSibling;

    if (!nextVideo) {
      return;
    }
    const nextVideoId = nextVideo.getAttribute("id");

    player.loadVideoById(nextVideoId);
    highlightSelectedItem(nextVideo);
  }
}

export function setPlayingVideo(videoId) {
  if (player) {
    player.loadVideoById(videoId);
  } else {
    player = new YT.Player("video-container", {
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
