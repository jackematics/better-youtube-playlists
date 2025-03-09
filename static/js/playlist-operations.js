const previous = document.getElementById("previous");

function handlePreviousClick() {
  const prevVideo = document
    .getElementById("playlist-items")
    .querySelector(".bg-warm-orange").previousElementSibling;

  if (prevVideo) {
    prevVideo.click();
  }
}

previous.addEventListener("click", handlePreviousClick);
