package server

import (
	"context"
	pb "github.com/Krokin/gocloudcamp/Part_Two/proto"
	log "github.com/Krokin/gocloudcamp/Part_Two/server/logger"
	"github.com/Krokin/gocloudcamp/Part_Two/server/playlist"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/durationpb"
	"reflect"
	"testing"
	"time"
)

var (
	EmptyPlaylist = func() *playlist.SongsPlaylist {
		return playlist.NewPlaylist()
	}
	Playlist = func() *playlist.SongsPlaylist {
		res := playlist.NewPlaylist()
		res.CreateSong("a", "s", time.Second*100)
		res.CreateSong("b", "s", time.Second*110)
		res.CreateSong("c", "s", time.Second*120)
		res.Play()
		return res
	}
	ctxWithError = ContextWithError()
)

func ContextWithError() context.Context {
	res, _ := context.WithTimeout(context.Background(), time.Nanosecond)
	return res
}

func TestServer_Create(t *testing.T) {
	type fields struct {
		Logger         *log.Logger
		P              playlist.Playlister
		PlaylistServer pb.PlaylistServer
	}
	type args struct {
		ctx context.Context
		req *pb.CreateSongRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *empty.Empty
		wantErr bool
	}{
		{
			"test grpc create empty playlist ok",
			fields{P: EmptyPlaylist()},
			args{context.TODO(), &pb.CreateSongRequest{
				Song: &pb.Song{Author: "sss", Title: "bbb", Dur: durationpb.New(time.Second * 200)},
			}},
			&empty.Empty{},
			false,
		}, {
			"test grpc update context err",
			fields{P: EmptyPlaylist()},
			args{ctxWithError, &pb.CreateSongRequest{}},
			nil,
			true,
		}, {
			"test grpc update not valid err",
			fields{P: EmptyPlaylist()},
			args{ctxWithError, &pb.CreateSongRequest{}},
			nil,
			true,
		}, {
			"test grpc create ok",
			fields{P: Playlist()},
			args{context.TODO(), &pb.CreateSongRequest{
				Song: &pb.Song{Author: "sss", Title: "bbb", Dur: durationpb.New(time.Second * 200)},
			}},
			&empty.Empty{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Logger:         tt.fields.Logger,
				P:              tt.fields.P,
				PlaylistServer: tt.fields.PlaylistServer,
			}
			got, err := s.Create(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_Delete(t *testing.T) {
	type fields struct {
		Logger         *log.Logger
		P              playlist.Playlister
		PlaylistServer pb.PlaylistServer
	}
	type args struct {
		ctx context.Context
		req *pb.DeleteSongRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *empty.Empty
		wantErr bool
	}{
		{
			"test grpc delete empty playlist",
			fields{P: EmptyPlaylist()},
			args{context.TODO(), &pb.DeleteSongRequest{}},
			nil,
			true,
		}, {
			"test grpc delete context err",
			fields{P: EmptyPlaylist()},
			args{ctxWithError, &pb.DeleteSongRequest{}},
			nil,
			true,
		}, {
			"test grpc delete not found",
			fields{P: Playlist()},
			args{context.TODO(), &pb.DeleteSongRequest{Author: "", Title: ""}},
			nil,
			true,
		}, {
			"test grpc delete play song err",
			fields{P: Playlist()},
			args{context.TODO(), &pb.DeleteSongRequest{Author: "s", Title: "a"}},
			nil,
			true,
		}, {
			"test grpc delete ok",
			fields{P: Playlist()},
			args{context.TODO(), &pb.DeleteSongRequest{Author: "s", Title: "b"}},
			&empty.Empty{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Logger:         tt.fields.Logger,
				P:              tt.fields.P,
				PlaylistServer: tt.fields.PlaylistServer,
			}
			got, err := s.Delete(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Delete() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_Next(t *testing.T) {
	type fields struct {
		Logger         *log.Logger
		P              playlist.Playlister
		PlaylistServer pb.PlaylistServer
	}
	type args struct {
		ctx context.Context
		e   *empty.Empty
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *empty.Empty
		wantErr bool
	}{
		{
			"test grpc next empty playlist",
			fields{P: EmptyPlaylist()},
			args{context.TODO(), &empty.Empty{}},
			nil,
			true,
		}, {
			"test grpc next context err",
			fields{P: EmptyPlaylist()},
			args{ctxWithError, &empty.Empty{}},
			nil,
			true,
		}, {
			"test grpc next ok",
			fields{P: Playlist()},
			args{context.TODO(), &empty.Empty{}},
			&empty.Empty{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Logger:         tt.fields.Logger,
				P:              tt.fields.P,
				PlaylistServer: tt.fields.PlaylistServer,
			}
			got, err := s.Next(tt.args.ctx, tt.args.e)
			if (err != nil) != tt.wantErr {
				t.Errorf("Next() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Next() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_Pause(t *testing.T) {
	type fields struct {
		Logger         *log.Logger
		P              playlist.Playlister
		PlaylistServer pb.PlaylistServer
	}
	type args struct {
		ctx context.Context
		e   *empty.Empty
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *empty.Empty
		wantErr bool
	}{
		{
			"test grpc pause empty playlist",
			fields{P: EmptyPlaylist()},
			args{context.TODO(), &empty.Empty{}},
			nil,
			true,
		}, {
			"test grpc pause context err",
			fields{P: EmptyPlaylist()},
			args{ctxWithError, &empty.Empty{}},
			nil,
			true,
		}, {
			"test grpc pause ok",
			fields{P: Playlist()},
			args{context.TODO(), &empty.Empty{}},
			&empty.Empty{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Logger:         tt.fields.Logger,
				P:              tt.fields.P,
				PlaylistServer: tt.fields.PlaylistServer,
			}
			got, err := s.Pause(tt.args.ctx, tt.args.e)
			if (err != nil) != tt.wantErr {
				t.Errorf("Pause() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pause() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_Play(t *testing.T) {
	type fields struct {
		Logger         *log.Logger
		P              playlist.Playlister
		PlaylistServer pb.PlaylistServer
	}
	type args struct {
		ctx context.Context
		e   *empty.Empty
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *empty.Empty
		wantErr bool
	}{
		{
			"test grpc play empty playlist",
			fields{P: EmptyPlaylist()},
			args{context.TODO(), &empty.Empty{}},
			nil,
			true,
		}, {
			"test grpc prev context err",
			fields{P: EmptyPlaylist()},
			args{ctxWithError, &empty.Empty{}},
			nil,
			true,
		}, {
			"test grpc prev ok",
			fields{P: Playlist()},
			args{context.TODO(), &empty.Empty{}},
			&empty.Empty{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Logger:         tt.fields.Logger,
				P:              tt.fields.P,
				PlaylistServer: tt.fields.PlaylistServer,
			}
			got, err := s.Play(tt.args.ctx, tt.args.e)
			if (err != nil) != tt.wantErr {
				t.Errorf("Play() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Play() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_Prev(t *testing.T) {
	type fields struct {
		Logger         *log.Logger
		P              playlist.Playlister
		PlaylistServer pb.PlaylistServer
	}
	type args struct {
		ctx context.Context
		e   *empty.Empty
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *empty.Empty
		wantErr bool
	}{
		{
			"test grpc prev empty playlist",
			fields{P: EmptyPlaylist()},
			args{context.TODO(), &empty.Empty{}},
			nil,
			true,
		}, {
			"test grpc prev context err",
			fields{P: EmptyPlaylist()},
			args{ctxWithError, &empty.Empty{}},
			nil,
			true,
		}, {
			"test grpc prev ok",
			fields{P: Playlist()},
			args{context.TODO(), &empty.Empty{}},
			&empty.Empty{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Logger:         tt.fields.Logger,
				P:              tt.fields.P,
				PlaylistServer: tt.fields.PlaylistServer,
			}
			got, err := s.Prev(tt.args.ctx, tt.args.e)
			if (err != nil) != tt.wantErr {
				t.Errorf("Prev() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Prev() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_ReadPlaylist(t *testing.T) {
	type fields struct {
		Logger         *log.Logger
		P              playlist.Playlister
		PlaylistServer pb.PlaylistServer
	}
	type args struct {
		ctx context.Context
		e   *empty.Empty
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.ReadPlaylistResponse
		wantErr bool
	}{
		{
			"test grpc readplaylist empty playlist",
			fields{P: EmptyPlaylist()},
			args{context.TODO(), &empty.Empty{}},
			nil,
			true,
		}, {
			"test grpc readplaylist context err",
			fields{P: EmptyPlaylist()},
			args{ctxWithError, &empty.Empty{}},
			nil,
			true,
		}, {
			"test grpc readplaylist ok",
			fields{P: Playlist()},
			args{context.TODO(), &empty.Empty{}},
			&pb.ReadPlaylistResponse{Songs: []*pb.Song{
				{Author: "s", Title: "a", Dur: durationpb.New(time.Second * 100)},
				{Author: "s", Title: "b", Dur: durationpb.New(time.Second * 110)},
				{Author: "s", Title: "c", Dur: durationpb.New(time.Second * 120)},
			}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Logger:         tt.fields.Logger,
				P:              tt.fields.P,
				PlaylistServer: tt.fields.PlaylistServer,
			}
			got, err := s.ReadPlaylist(tt.args.ctx, tt.args.e)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadPlaylist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadPlaylist() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_ReadSong(t *testing.T) {
	type fields struct {
		Logger         *log.Logger
		P              playlist.Playlister
		PlaylistServer pb.PlaylistServer
	}
	type args struct {
		ctx context.Context
		e   *empty.Empty
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.ReadSongResponse
		wantErr bool
	}{
		{
			"test grpc readsong empty playlist",
			fields{P: EmptyPlaylist()},
			args{context.TODO(), &empty.Empty{}},
			nil,
			true,
		}, {
			"test grpc readsong context err",
			fields{P: EmptyPlaylist()},
			args{ctxWithError, &empty.Empty{}},
			nil,
			true,
		}, {
			"test grpc readsong ok",
			fields{P: Playlist()},
			args{context.TODO(), &empty.Empty{}},
			&pb.ReadSongResponse{Song: &pb.Song{Author: "s", Title: "a", Dur: durationpb.New(time.Second * 100)}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Logger:         tt.fields.Logger,
				P:              tt.fields.P,
				PlaylistServer: tt.fields.PlaylistServer,
			}
			got, err := s.ReadSong(tt.args.ctx, tt.args.e)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadSong() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadSong() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_Update(t *testing.T) {
	type fields struct {
		Logger         *log.Logger
		P              playlist.Playlister
		PlaylistServer pb.PlaylistServer
	}
	type args struct {
		ctx context.Context
		req *pb.UpdateSongRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *empty.Empty
		wantErr bool
	}{
		{
			"test grpc update empty playlist",
			fields{P: EmptyPlaylist()},
			args{context.TODO(), &pb.UpdateSongRequest{
				Author: "s",
				Title:  "b",
				Song:   &pb.Song{Author: "sss", Title: "bbb", Dur: durationpb.New(time.Second * 200)},
			}},
			nil,
			true,
		}, {
			"test grpc update context err",
			fields{P: EmptyPlaylist()},
			args{ctxWithError, &pb.UpdateSongRequest{}},
			nil,
			true,
		}, {
			"test grpc update not valid err",
			fields{P: EmptyPlaylist()},
			args{ctxWithError, &pb.UpdateSongRequest{}},
			nil,
			true,
		}, {
			"test grpc update ok",
			fields{P: Playlist()},
			args{context.TODO(), &pb.UpdateSongRequest{
				Author: "s",
				Title:  "b",
				Song:   &pb.Song{Author: "sss", Title: "bbb", Dur: durationpb.New(time.Second * 200)},
			}},
			&empty.Empty{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Logger:         tt.fields.Logger,
				P:              tt.fields.P,
				PlaylistServer: tt.fields.PlaylistServer,
			}
			got, err := s.Update(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update() got = %v, want %v", got, tt.want)
			}
		})
	}
}
