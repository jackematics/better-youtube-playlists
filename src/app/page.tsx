"use client";

import { useState } from "react";
import AddPlaylistModal from "./components/playlist-list/AddPlaylistModal";
import PlaylistList from "./components/playlist-list/PlaylistList";
import useYoutubeMetadataFetcher from "./hooks/useYoutubeMetadataFetcher";
import {
  PlaylistMetadata,
  SelectedPlaylistMetadata,
} from "./types/youtube-playlist-metadata-types";
import PlaylistDescription from "./components/playlist-description/PlaylistDescription";
import useYoutubePlaylistItemsFetcher from "./hooks/useYoutubePlaylistItemsFetcher";
import {
  PlaylistData,
  PlaylistItem,
} from "./types/youtube-playlist-items-types";
import Playlist from "./components/playlist/Playlist";
import YoutubeVideoEmbed from "./components/youtube-video-embed/YoutubeVideoEmbed";
import PlaylistOperations from "./components/playlist-operations/PlaylistOperations";

const Home = () => {
  const [addPlaylistModalOpen, setAddPlaylistModalOpen] =
    useState<boolean>(false);
  const {
    updatePlaylistMetadataState,
    playlistMetadataCollection,
    playlistMetadataValidationResult,
  } = useYoutubeMetadataFetcher();

  const { updatePlaylistItemCollectionState, playlistItemCollection } =
    useYoutubePlaylistItemsFetcher();

  const [selectedPlaylistMetadata, setSelectedPlaylistMetadata] = useState<
    SelectedPlaylistMetadata | undefined
  >();
  const [selectedPlaylistData, setSelectedPlaylistData] = useState<
    PlaylistData | undefined
  >();
  const [selectedPlaylistItem, setSelectedPlaylistItem] = useState<
    PlaylistItem | undefined
  >();
  const [currentVideoIndex, setCurrentVideoIndex] = useState<number>(0);
  const [playNextVideo, setPlayNextVideo] = useState<boolean>(false);
  const [playPrevVideo, setPlayPrevVideo] = useState<boolean>(false);

  const openAddPlaylistModal = (): void => {
    setAddPlaylistModalOpen(true);
  };

  const handleAddPlaylistId = async (playlistId: string): Promise<void> => {
    await updatePlaylistMetadataState(playlistId);

    if (playlistMetadataValidationResult.valid) {
      await updatePlaylistItemCollectionState(playlistId);
    }
  };

  const closeAddPlaylistModal = (): void => {
    if (addPlaylistModalOpen) {
      setAddPlaylistModalOpen(false);
    }
  };

  const handleSelectPlaylist = (playlistMetadata: PlaylistMetadata): void => {
    const selectedPlaylist = playlistItemCollection.find(
      (item: PlaylistData) => item.id === playlistMetadata.id
    )!;

    setSelectedPlaylistMetadata({
      ...playlistMetadata,
      totalResults: selectedPlaylist.totalResults,
    });
    setSelectedPlaylistData(selectedPlaylist);
  };

  const handlePlaylistItemSelect = (
    playlistItem: PlaylistItem,
    itemIndex: number
  ) => {
    setSelectedPlaylistItem(playlistItem);
    setCurrentVideoIndex(itemIndex);
    setPlayNextVideo(false);
    setPlayPrevVideo(false);
  };

  const handlePlayNextVideo = () => {
    setPlayNextVideo(true);
  };

  const handlePlayPrevVideo = () => {
    setPlayPrevVideo(true);
  };

  return (
    <main className="bg-background-dark-blue">
      <div
        className={`flex justify-center ${
          addPlaylistModalOpen ? "modal-backdrop pointer-events-none" : ""
        }`}
      >
        <PlaylistList
          playlistMetadataCollection={playlistMetadataCollection}
          openAddPlaylistModalCallback={openAddPlaylistModal}
          selectPlaylistCallback={handleSelectPlaylist}
        />
        <div>
          <div className="flex flex-row">
            <YoutubeVideoEmbed
              selectedPlaylistItem={selectedPlaylistItem}
              videoEndCallback={handlePlayNextVideo}
            />
            <div className="flex-initial w-[40rem] h-[26rem] min-w-[30rem] bg-container-dark-blue mt-4 border-4 relative">
              <PlaylistDescription
                selectedPlaylistMetadata={selectedPlaylistMetadata}
                currentVideoIndex={currentVideoIndex}
              />
              <div className="pt-10 pl-5">
                {selectedPlaylistMetadata && (
                  <PlaylistOperations
                    data-testid="playlist-operations"
                    prevPlaylistItemCallback={handlePlayPrevVideo}
                    nextPlaylistItemCallback={handlePlayNextVideo}
                  />
                )}
              </div>
            </div>
          </div>
          <Playlist
            selectedPlaylistData={selectedPlaylistData}
            nextVideoSelected={playNextVideo}
            prevVideoSelected={playPrevVideo}
            playlistItemSelectCallback={handlePlaylistItemSelect}
          />
        </div>
      </div>
      <div
        className={addPlaylistModalOpen ? "" : "hidden"}
        data-testid="add-playlist-modal-wrapper"
      >
        <AddPlaylistModal
          playlistMetadataValidationResult={playlistMetadataValidationResult}
          addPlaylistIdCallback={handleAddPlaylistId}
          closePlaylistModalCallback={closeAddPlaylistModal}
        />
      </div>
    </main>
  );
};

export default Home;
