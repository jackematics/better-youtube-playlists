import { handleNextClick } from "./playlist-operations.js";

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
