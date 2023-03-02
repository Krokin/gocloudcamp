package errors

import (
	"errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrorEmptyPlaylist = errors.New("playlist empty")
	ErrorPlaySongNow   = errors.New("the song is playing now")
	ErrorNotValid      = errors.New("invalid values")
	ErrorNotFound      = errors.New("songs not in playlist")
)

func Error(err error) error {
	var newError error
	switch err {
	case context.DeadlineExceeded, context.Canceled:
		newError = status.Error(codes.Canceled, err.Error())
	case ErrorEmptyPlaylist, ErrorPlaySongNow:
		newError = status.Error(codes.FailedPrecondition, err.Error())
	case ErrorNotFound:
		newError = status.Error(codes.NotFound, err.Error())
	case ErrorNotValid:
		newError = status.Error(codes.InvalidArgument, err.Error())
	default:
		newError = status.Error(codes.Unknown, err.Error())
	}
	return newError
}
