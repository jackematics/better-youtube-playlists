function clearModal() {
  document.getElementById("playlist-id-input").value = "";
  document.getElementById("validation-message").textContent = "";
}

document
  .getElementById("open-add-playlist-modal-btn")
  .addEventListener("click", function () {
    document.getElementById("add-playlist-modal").classList.remove("invisible");
  });

document
  .getElementById("close-add-playlist-modal-btn")
  .addEventListener("click", function () {
    document.getElementById("add-playlist-modal").classList.add("invisible");
    clearModal();
  });
