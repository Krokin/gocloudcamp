package Part_One

import "errors"

var (
	ErrorEmptyPlaylist = errors.New("playlist empty")
	ErrorPlaySongNow   = errors.New("the song is playing now")
	ErrorNotValid      = errors.New("invalid values")
	ErrorNotFound      = errors.New("songs not in playlist")
)
