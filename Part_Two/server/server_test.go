package server

//
//import (
//	"context"
//	sourse "github.com/Krokin/gocloudcamp/Part_Two/proto"
//	"github.com/Krokin/gocloudcamp/Part_Two/server/playlist"
//	"github.com/golang/protobuf/ptypes/empty"
//	"reflect"
//	"testing"
//)
//
//func TestNewServer(t *testing.T) {
//	tests := []struct {
//		name string
//		want *Server
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := NewServer(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("NewServer() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestServer_Create(t *testing.T) {
//	type fields struct {
//		Logger         *log.Logger
//		P              playlist.Playlister
//		PlaylistServer sourse.PlaylistServer
//	}
//	type args struct {
//		ctx context.Context
//		req *pb.CreateSongRequest
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    *empty.Empty
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			s := &Server{
//				Logger:         tt.fields.Logger,
//				P:              tt.fields.P,
//				PlaylistServer: tt.fields.PlaylistServer,
//			}
//			got, err := s.Create(tt.args.ctx, tt.args.req)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Create() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestServer_Delete(t *testing.T) {
//	type fields struct {
//		Logger         *log.Logger
//		P              playlist.Playlister
//		PlaylistServer sourse.PlaylistServer
//	}
//	type args struct {
//		ctx context.Context
//		req *pb.DeleteSongRequest
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    *empty.Empty
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			s := &Server{
//				Logger:         tt.fields.Logger,
//				P:              tt.fields.P,
//				PlaylistServer: tt.fields.PlaylistServer,
//			}
//			got, err := s.Delete(tt.args.ctx, tt.args.req)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Delete() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestServer_Next(t *testing.T) {
//	type fields struct {
//		Logger         *log.Logger
//		P              playlist.Playlister
//		PlaylistServer sourse.PlaylistServer
//	}
//	type args struct {
//		ctx context.Context
//		e   *empty.Empty
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    *empty.Empty
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			s := &Server{
//				Logger:         tt.fields.Logger,
//				P:              tt.fields.P,
//				PlaylistServer: tt.fields.PlaylistServer,
//			}
//			got, err := s.Next(tt.args.ctx, tt.args.e)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("Next() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Next() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestServer_Pause(t *testing.T) {
//	type fields struct {
//		Logger         *log.Logger
//		P              playlist.Playlister
//		PlaylistServer sourse.PlaylistServer
//	}
//	type args struct {
//		ctx context.Context
//		e   *empty.Empty
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    *empty.Empty
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			s := &Server{
//				Logger:         tt.fields.Logger,
//				P:              tt.fields.P,
//				PlaylistServer: tt.fields.PlaylistServer,
//			}
//			got, err := s.Pause(tt.args.ctx, tt.args.e)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("Pause() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Pause() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestServer_Play(t *testing.T) {
//	type fields struct {
//		Logger         *log.Logger
//		P              playlist.Playlister
//		PlaylistServer sourse.PlaylistServer
//	}
//	type args struct {
//		ctx context.Context
//		e   *empty.Empty
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    *empty.Empty
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			s := &Server{
//				Logger:         tt.fields.Logger,
//				P:              tt.fields.P,
//				PlaylistServer: tt.fields.PlaylistServer,
//			}
//			got, err := s.Play(tt.args.ctx, tt.args.e)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("Play() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Play() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestServer_Prev(t *testing.T) {
//	type fields struct {
//		Logger         *log.Logger
//		P              playlist.Playlister
//		PlaylistServer sourse.PlaylistServer
//	}
//	type args struct {
//		ctx context.Context
//		e   *empty.Empty
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    *empty.Empty
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			s := &Server{
//				Logger:         tt.fields.Logger,
//				P:              tt.fields.P,
//				PlaylistServer: tt.fields.PlaylistServer,
//			}
//			got, err := s.Prev(tt.args.ctx, tt.args.e)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("Prev() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Prev() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestServer_ReadPlaylist(t *testing.T) {
//	type fields struct {
//		Logger         *log.Logger
//		P              playlist.Playlister
//		PlaylistServer sourse.PlaylistServer
//	}
//	type args struct {
//		ctx context.Context
//		e   *empty.Empty
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    *pb.ReadPlaylistResponse
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			s := &Server{
//				Logger:         tt.fields.Logger,
//				P:              tt.fields.P,
//				PlaylistServer: tt.fields.PlaylistServer,
//			}
//			got, err := s.ReadPlaylist(tt.args.ctx, tt.args.e)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("ReadPlaylist() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("ReadPlaylist() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestServer_ReadSong(t *testing.T) {
//	type fields struct {
//		Logger         *log.Logger
//		P              playlist.Playlister
//		PlaylistServer sourse.PlaylistServer
//	}
//	type args struct {
//		ctx context.Context
//		e   *empty.Empty
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    *pb.ReadSongResponse
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			s := &Server{
//				Logger:         tt.fields.Logger,
//				P:              tt.fields.P,
//				PlaylistServer: tt.fields.PlaylistServer,
//			}
//			got, err := s.ReadSong(tt.args.ctx, tt.args.e)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("ReadSong() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("ReadSong() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestServer_Update(t *testing.T) {
//	type fields struct {
//		Logger         *log.Logger
//		P              playlist.Playlister
//		PlaylistServer sourse.PlaylistServer
//	}
//	type args struct {
//		ctx context.Context
//		req *pb.UpdateSongRequest
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    *empty.Empty
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			s := &Server{
//				Logger:         tt.fields.Logger,
//				P:              tt.fields.P,
//				PlaylistServer: tt.fields.PlaylistServer,
//			}
//			got, err := s.Update(tt.args.ctx, tt.args.req)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Update() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
