# better-youtube-playlists

An improved user experience for anyone who likes to listen to music using Youtube playlists but finds Youtube's current implementation lacking.

## To Do

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
  - Should populate playlist items
- Cache: should cache playlist list data so users returning to the site will automatically have playlists loaded
- Complete metadata fetch validation
- Option to delete playlists

### Playlist Items

- Items UI
- Default: Empty container
- List playlist items of selected playlist
  - Item number, thumbnail, video title
- Scroll through items
- Selecting an item highlights it
- Playlist titles beyond a certain character length move to the next line
- Selecting an item scrolls that item to the middle of the container
- Validate items fetch
- Load more than 50 items on playlist select

### Video Player

- Default: show a graphic
- On playlist select: Play the first video in the list
- On playlist item select: display and play video corresponding to selected item
- On video end: move to the next playlist item which also focuses it (Not sure how to test so didn't) :innocent:
- On playlist end: stop

### Playlist Description

- :white_check_mark: ~~Default: title only, No Playlist Selected~~
- :white_check_mark: ~~On playlist selection: Show title~~
- :white_check_mark: ~~On playlist selection: Show title, owner and number of videos~~
- Show the index of the current video in the description

### Playlist Operations

- Previous: Move to the previous playlist item
- Next: Move to the next playlist item
- Shuffle: At the end of every video
  - Randomly select a new playlist item at the end of every video
  - Move list scroll down to the selected video
- Loop: At the end of the playlist, loop back to the first video in the playlist

### Contact

- Github link
- Email

### Misc

- Description and Playlist Items shrinks on lower screen sizes
- Fix skipping through playlist quickly bug

### Stretch Goals

- Mobile UI
- Contact details
- Ads
- Merged playlists
- Keyboard shortcuts if possible
