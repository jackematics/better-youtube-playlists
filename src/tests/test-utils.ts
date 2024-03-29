import { fireEvent, screen, act } from "@testing-library/react";

const addTestPlaylistPath = async () => {
  const addPlaylistButton = screen.getByRole("button", {
    name: /Add Playlist/i,
  });
  const playlistIdInput = screen.getByTestId("playlist-id-input");
  const addButton = screen.getByTestId("add-id-button");

  fireEvent.click(addPlaylistButton);
  fireEvent.change(playlistIdInput, {
    target: { value: "test-playlist-id" },
  });
  fireEvent.click(addButton);
};

export { addTestPlaylistPath };
