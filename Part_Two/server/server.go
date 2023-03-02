package server

import (
	"context"

	pb "main/Part_Two/proto"
	"main/Part_Two/server/errors"
	log "main/Part_Two/server/logger"
	pl "main/Part_Two/server/playlist"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/durationpb"
)

type Server struct {
	Logger *log.Logger
	P      pl.Playlister
	pb.PlaylistServer
}

func NewServer() *Server {
	logger := log.NewLogger()
	playlist := pl.NewPlaylist()
	if err := playlist.LoadPlaylist(); err != nil {
		logger.Log.Println("Playlist created")
	} else {
		logger.Log.Println("Playlist load")
	}
	return &Server{P: playlist, Logger: logger}
}

func (s *Server) Play(ctx context.Context, e *empty.Empty) (*empty.Empty, error) {
	err := ctx.Err()
	if err != nil {
		return nil, errors.Error(err)
	}
	err = s.P.Play()
	if err != nil {
		return nil, errors.Error(err)
	}
	return &empty.Empty{}, nil
}

func (s *Server) Pause(ctx context.Context, e *empty.Empty) (*empty.Empty, error) {
	err := ctx.Err()
	if err != nil {
		return nil, errors.Error(err)
	}
	err = s.P.Pause()
	if err != nil {
		return nil, errors.Error(err)
	}
	return &empty.Empty{}, nil
}
func (s *Server) Next(ctx context.Context, e *empty.Empty) (*empty.Empty, error) {
	err := ctx.Err()
	if err != nil {
		return nil, errors.Error(err)
	}
	err = s.P.Next()
	if err != nil {
		return nil, errors.Error(err)
	}
	return &empty.Empty{}, nil
}
func (s *Server) Prev(ctx context.Context, e *empty.Empty) (*empty.Empty, error) {
	err := ctx.Err()
	if err != nil {
		return nil, errors.Error(err)
	}
	err = s.P.Prev()
	if err != nil {
		return nil, errors.Error(err)
	}
	return &empty.Empty{}, nil
}
func (s *Server) Create(ctx context.Context, req *pb.CreateSongRequest) (*empty.Empty, error) {
	err := ctx.Err()
	if err != nil {
		return nil, errors.Error(err)
	}
	err = s.P.CreateSong(req.GetSong().GetTitle(), req.GetSong().GetAuthor(), req.GetSong().GetDur().AsDuration())
	if err != nil {
		return nil, errors.Error(err)
	}
	return &empty.Empty{}, nil
}

func (s *Server) Delete(ctx context.Context, req *pb.DeleteSongRequest) (*empty.Empty, error) {
	err := ctx.Err()
	if err != nil {
		return nil, errors.Error(err)
	}
	err = s.P.DeleteSong(req.GetTitle(), req.GetAuthor())
	if err != nil {
		if err == errors.ErrorPlaySongNow {
			return nil, errors.Error(err)
		}
		return nil, errors.Error(err)
	}
	return &empty.Empty{}, nil
}

func (s *Server) ReadSong(ctx context.Context, e *empty.Empty) (*pb.ReadSongResponse, error) {
	err := ctx.Err()
	if err != nil {
		return nil, errors.Error(err)
	}
	r, err := s.P.ReadSong()
	if err != nil {
		return nil, errors.Error(err)
	}
	return &pb.ReadSongResponse{Song: &pb.Song{Author: r.Author, Title: r.Title, Dur: durationpb.New(r.Duration)}}, nil
}

func (s *Server) ReadPlaylist(ctx context.Context, e *empty.Empty) (*pb.ReadPlaylistResponse, error) {
	err := ctx.Err()
	if err != nil {
		return nil, errors.Error(err)
	}
	r, err := s.P.ReadPlaylist()
	if err != nil {
		return nil, errors.Error(err)
	}
	res := make([]*pb.Song, 0, len(r))
	for _, song := range r {
		res = append(res, &pb.Song{Author: song.Author, Title: song.Title, Dur: durationpb.New(song.Duration)})
	}
	return &pb.ReadPlaylistResponse{Songs: res}, nil
}

func (s *Server) Update(ctx context.Context, req *pb.UpdateSongRequest) (*empty.Empty, error) {
	err := ctx.Err()
	if err != nil {
		return nil, errors.Error(err)
	}
	song, err := pl.NewSong(req.GetSong().GetTitle(), req.GetSong().GetAuthor(), req.GetSong().GetDur().AsDuration())
	if err != nil {
		return nil, errors.Error(err)
	}
	err = s.P.UpdateInfo(req.GetTitle(), req.GetAuthor(), *song)
	if err != nil {
		return nil, errors.Error(err)
	}
	return &empty.Empty{}, nil
}
