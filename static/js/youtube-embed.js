export function setPlayingVideo(id, title) {
  const embed = document.getElementById("youtube-embed");
  embed.innerHTML = `
          <div
            class="w-full h-full bg-black grid place-items-center relative"
          >
            <iframe
              class="w-full h-full"
              src="https://www.youtube.com/embed/${id}?autoplay=1&rel=0"
              title="${title}"
              allowfullscreen
              allow="autoplay"
              frameborder="0"
            />
          </div>
  `;
}
