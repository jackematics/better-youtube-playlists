CREATE TABLE IF NOT EXISTS user (
    id TEXT PRIMARY KEY,
    modal_hidden BOOLEAN DEFAULT TRUE,
    modal_validation_message TEXT,
    last_updated_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS playlist (
    id TEXT PRIMARY KEY,
    title TEXT,
    channel_owner TEXT,
    total_videos INT
);

CREATE TABLE IF NOT EXISTS user_playlist (
    user_id TEXT,
    playlist_id TEXT,
    selected BOOLEAN DEFAULT FALSE
    CONSTRAINT fk_user
        FOREIGN KEY (user_id)
            REFERENCES user (id)
    CONSTRAINT fk_playlist
        FOREIGN KEY (playlist_id)
            REFERENCES playlist (id)
);

CREATE TABLE IF NOT EXISTS playlist_item (
    id TEXT PRIMARY KEY,
    title TEXT,
    thumbnail_url TEXT
);

CREATE TABLE IF NOT EXISTS playlist_playlist_item (
    id TEXT,
    playlist_id TEXT,
    playlist_item_id TEXT
    CONSTRAINT fk_playlist
        FOREIGN KEY (playlist_id)
            REFERENCES playlist (id)
    CONSTRAINT fk_playlist_item
        FOREIGN KEY (playlist_item_id)
            REFERENCES playlist_item (id)
);

CREATE TABLE IF NOT EXISTS user_playlist_playlist_item (
    user_id TEXT,
    playlist_playlist_item_id TEXT,
    selected BOOLEAN DEFAULT FALSE
    CONSTRAINT fk_user
        FOREIGN KEY (user_id)
            REFERENCES user (id)
    CONSTRAINT fk_playlist_playlist_item
        FOREIGN KEY (playlist_playlist_item_id)
            REFERENCES playlist_playlist_item (id)
);