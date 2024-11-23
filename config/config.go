package config

import "os"

type EnvConfig struct {
	YoutubeApiKey string
}

var Config = EnvConfig {
	YoutubeApiKey: os.Getenv("YOUTUBE_API_KEY"),
}