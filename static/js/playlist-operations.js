function handlePreviousClick() {
  const prevVideo = document
    .getElementById("playlist-items")
    .querySelector(".bg-warm-orange").previousElementSibling;

  if (prevVideo) {
    prevVideo.click();
  }
}

function handleNextClick() {
  const nextVideo = document
    .getElementById("playlist-items")
    .querySelector(".bg-warm-orange").nextElementSibling;

  if (nextVideo) {
    nextVideo.click();
  }
}

const previous = document.getElementById("previous");
previous.addEventListener("click", handlePreviousClick);

const next = document.getElementById("next");
next.addEventListener("click", handleNextClick);
