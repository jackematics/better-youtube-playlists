export function highlightSelectedItem(selected) {
  document.querySelectorAll(".playlist-items").forEach((item) => {
    item.classList.remove("bg-warm-orange");
  });

  selected.classList.add("bg-warm-orange");
}
