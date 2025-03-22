const addPlaylistModalEl = document.getElementById("add-playlist-modal");
const playlistIdInputEl = document.getElementById("playlist-id-input");
const modalValidationMessageEl = document.getElementById(
  "modal-validation-message"
);
const openPlaylistModalBtnEl = document.getElementById(
  "open-add-playlist-modal-btn"
);
const closeAddPlaylistModalBtnEl = document.getElementById(
  "close-add-playlist-modal-btn"
);

export function closeModal() {
  addPlaylistModalEl.classList.add("invisible");
  clearModal();
}

function clearModal() {
  playlistIdInputEl("playlist-id-input").value = "";
  modalValidationMessageEl("modal-validation-message").textContent = "";
}

openPlaylistModalBtnEl.addEventListener("click", function () {
  addPlaylistModalEl.classList.remove("invisible");
});

closeAddPlaylistModalBtnEl.addEventListener("click", function () {
  closeModal();
});
