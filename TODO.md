# TODO

### Playlist List

- :white_check_mark: ~~PlaylistList UI (Desktop)~~
- :white_check_mark: ~~Add Playlist Modal UI (Desktop)~~
- :white_check_mark: ~~Add playlist~~
  - :white_check_mark: ~~On click, should raise a modal asking for the list url of a youtube playlist~~
  - :white_check_mark: ~~Selecting cancel button should close the modal~~
  - :white_check_mark: ~~Adding a list id and selecting the add button should add a playlist to the list of playlists and close the modal~~
  - :white_check_mark: ~~Attempting to add an empty playlist id should show an error validation message~~
  - :white_check_mark: ~~Attempting to add an invalid playist id should show an error validation message~~
  - :white_check_mark: ~~Attempting to add a duplicate playlist id should show an error validation message~~
  - :white_check_mark: ~~Elements outside of the modal while the modal is open should be unclickable~~
  - :white_check_mark: ~~Playlist id input and validation messages should be cleared if the playlist is closed~~
  - :white_check_mark: ~~Log handle playlist quota exceeded (code 403) and return validation message to user~~
  - :white_check_mark: ~~Log error with youtube api and return validation message to user~~
  - :white_check_mark: ~~Log error with youtube api key and return validation message to user~~
- Select playlist
  - :white_check_mark: ~~Should populate the playlist description~~
  - :white_check_mark: ~~Should highlight the selected playlist~~
  - :white_check_mark: ~~Should populate playlist items~~
  - Should show a toast if there's an error fetching playlist items
- :white_check_mark: ~~Cache: should cache playlist list data so users returning to the site will automatically have playlists loaded~~
- Handle empty playlists
- Selecting a playlist resets the playlist operations and playlist history

### Playlist Items

- :white_check_mark: ~~Items UI~~
- :white_check_mark: ~~Default: Empty container~~
- :white_check_mark: ~~List playlist items of selected playlist~~
  - :white_check_mark: ~~Item number, thumbnail, video title~~
- :white_check_mark: ~~Scroll through items~~
- :white_check_mark: ~~Selecting an item highlights it~~
- :white_check_mark: ~~Playlist titles beyond a certain character length are truncated~~
- :white_check_mark: ~~Selecting an item scrolls that item to the middle of the container~~
- :white_check_mark: ~~Validate items fetch~~
- :white_check_mark: ~~Load more than 50 items on playlist select~~
- :white_check_mark: ~~Loading spinner when loading items~~
- :white_check_mark: ~~Selecting the same playlist multiple times won't refetch~~
- Handle unavailable videos
- Handle empty playlists
- Handle private videos
- Handle age restricted videos

### Video Player

- :white_check_mark: ~~Default: show a graphic~~
- :white_check_mark: ~~On playlist item select: display and play video corresponding to selected item~~
- :white_check_mark: ~~On playlist select: Play the first video in the list~~
- :white_check_mark: ~~On video end: move to the next playlist item which also focuses it~~
- :white_check_mark: ~~On playlist end: stop~~

### Playlist Description

- :white_check_mark: ~Default: title only, No Playlist Selected~
- :white_check_mark: ~On playlist selection: Show title~
- :white_check_mark: ~On playlist selection: Show title, owner and number of videos~
- :white_check_mark: ~Show the index of the current video in the description~
- Buy me a coffee button?

### Playlist Operations

- :white_check_mark: ~Previous: Move to the previous playlist item~
- :white_check_mark: ~Next: Move to the next playlist item~
- Random
  - :white_check_mark: ~At the end of every video~
    - :white_check_mark: ~Randomly select a new playlist item~
    - :white_check_mark: ~scroll down to the selected video~
  - :white_check_mark: ~Pressing the next button~
    - :white_check_mark: ~Randomly selects a new playlist item~
    - :white_check_mark: ~scroll down to the selected video~
  - Pressing the previous button
    - Returns to the previously played video
    - Returns to the previously indexed video if the history is empty
  - Deselecting
    - Resets playlist history
- Shuffle
  - Deselecting
    - Resets playist history
- Loop: At the end of the playlist, loop back to the first video in the playlist
- Bin: Remove playlists from list
- Download: Download a list of all the videos as MP3/MP4 files
- Clicking next / prev button shortcuts will move to the next / prev video even when not on the same tab

### Contact

- Github link
- Email

### Misc

- Make it accessible
- Use Google Fonts? Better fontage?

### Stretch Goals

- Mobile UI
- Contact details
- Ads
- Merged playlists
- Keyboard shortcuts if possible

### Refactoring

- Add more sad path tests
