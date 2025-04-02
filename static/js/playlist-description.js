const totalVideosEl = document.getElementById("total-videos");

let CURRENT_VIDEO_STATE = {
  currentIndex: null,
  totalVideos: null,
  unavailableVideoCount: null,
};

export const CurrentVideoState = {
  setState: (newState) => {
    CURRENT_VIDEO_STATE = { ...CURRENT_VIDEO_STATE, ...newState };
  },
  clear: () => {
    CURRENT_VIDEO_STATE = {
      currentIndex: null,
      totalVideos: null,
      unavailableVideoCount: null,
    };
  },
  getState: () => {
    return CURRENT_VIDEO_STATE;
  },
  render: () => {
    const { currentIndex, totalVideos, unavailableVideoCount } =
      CURRENT_VIDEO_STATE;
    totalVideosEl.textContent =
      currentIndex && totalVideos
        ? `Video: ${currentIndex} / ${totalVideos}`
        : "";

    if (currentIndex && totalVideos && unavailableVideoCount) {
      totalVideosEl.textContent += ` (${unavailableVideoCount} unavailable videos hidden)`;
    }
  },
};
