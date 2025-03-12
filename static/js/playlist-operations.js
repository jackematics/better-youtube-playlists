let RANDOMISE = false;

function handlePreviousClick() {
  const prevVideo = document
    .getElementById("playlist-items")
    .querySelector(".bg-warm-orange").previousElementSibling;

  if (prevVideo) {
    prevVideo.click();
  }
}

export function handleNextClick() {
  const currentVideo = document
    .getElementById("playlist-items")
    .querySelector(".bg-warm-orange");

  let nextVideo;
  if (RANDOMISE) {
    const currentItemIndex = Number.parseInt(
      currentVideo.children[0].children[0].textContent
    );
    const totalVideosSplit = document
      .getElementById("total-videos")
      .textContent.split(" ");
    const totalVideoCount = Number.parseInt(
      totalVideosSplit[totalVideosSplit.length - 1]
    );

    let nextLiIndex;
    do {
      nextLiIndex = Math.floor(Math.random() * totalVideoCount);
    } while (nextLiIndex === currentItemIndex);

    nextVideo = document.getElementById("playlist-items").children[nextLiIndex];
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
    event.currentTarget.classList.remove("bg-orange-highlight");
    event.currentTarget.classList.add("bg-white");
  }
}

document
  .getElementById("previous")
  .addEventListener("click", handlePreviousClick);

document.getElementById("next").addEventListener("click", handleNextClick);

document.getElementById("randomise").addEventListener("click", handleRandomise);
