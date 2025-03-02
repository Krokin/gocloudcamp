// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: Part_Two/proto/route.proto

package sourse

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PlaylistClient is the client API for Playlist service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PlaylistClient interface {
	Play(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	Pause(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	Next(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	Prev(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	Create(ctx context.Context, in *CreateSongRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Delete(ctx context.Context, in *DeleteSongRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	ReadSong(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ReadSongResponse, error)
	ReadPlaylist(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ReadPlaylistResponse, error)
	Update(ctx context.Context, in *UpdateSongRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type playlistClient struct {
	cc grpc.ClientConnInterface
}

func NewPlaylistClient(cc grpc.ClientConnInterface) PlaylistClient {
	return &playlistClient{cc}
}

func (c *playlistClient) Play(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/grpc.Playlist/Play", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playlistClient) Pause(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/grpc.Playlist/Pause", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playlistClient) Next(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/grpc.Playlist/Next", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playlistClient) Prev(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/grpc.Playlist/Prev", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playlistClient) Create(ctx context.Context, in *CreateSongRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/grpc.Playlist/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playlistClient) Delete(ctx context.Context, in *DeleteSongRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/grpc.Playlist/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playlistClient) ReadSong(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ReadSongResponse, error) {
	out := new(ReadSongResponse)
	err := c.cc.Invoke(ctx, "/grpc.Playlist/ReadSong", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playlistClient) ReadPlaylist(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ReadPlaylistResponse, error) {
	out := new(ReadPlaylistResponse)
	err := c.cc.Invoke(ctx, "/grpc.Playlist/ReadPlaylist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playlistClient) Update(ctx context.Context, in *UpdateSongRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/grpc.Playlist/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlaylistServer is the server API for Playlist service.
// All implementations must embed UnimplementedPlaylistServer
// for forward compatibility
type PlaylistServer interface {
	Play(context.Context, *empty.Empty) (*empty.Empty, error)
	Pause(context.Context, *empty.Empty) (*empty.Empty, error)
	Next(context.Context, *empty.Empty) (*empty.Empty, error)
	Prev(context.Context, *empty.Empty) (*empty.Empty, error)
	Create(context.Context, *CreateSongRequest) (*empty.Empty, error)
	Delete(context.Context, *DeleteSongRequest) (*empty.Empty, error)
	ReadSong(context.Context, *empty.Empty) (*ReadSongResponse, error)
	ReadPlaylist(context.Context, *empty.Empty) (*ReadPlaylistResponse, error)
	Update(context.Context, *UpdateSongRequest) (*empty.Empty, error)
	mustEmbedUnimplementedPlaylistServer()
}

// UnimplementedPlaylistServer must be embedded to have forward compatible implementations.
type UnimplementedPlaylistServer struct {
}

func (UnimplementedPlaylistServer) Play(context.Context, *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Play not implemented")
}
func (UnimplementedPlaylistServer) Pause(context.Context, *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pause not implemented")
}
func (UnimplementedPlaylistServer) Next(context.Context, *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Next not implemented")
}
func (UnimplementedPlaylistServer) Prev(context.Context, *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Prev not implemented")
}
func (UnimplementedPlaylistServer) Create(context.Context, *CreateSongRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPlaylistServer) Delete(context.Context, *DeleteSongRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPlaylistServer) ReadSong(context.Context, *empty.Empty) (*ReadSongResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadSong not implemented")
}
func (UnimplementedPlaylistServer) ReadPlaylist(context.Context, *empty.Empty) (*ReadPlaylistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadPlaylist not implemented")
}
func (UnimplementedPlaylistServer) Update(context.Context, *UpdateSongRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedPlaylistServer) mustEmbedUnimplementedPlaylistServer() {}

// UnsafePlaylistServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PlaylistServer will
// result in compilation errors.
type UnsafePlaylistServer interface {
	mustEmbedUnimplementedPlaylistServer()
}

func RegisterPlaylistServer(s grpc.ServiceRegistrar, srv PlaylistServer) {
	s.RegisterService(&Playlist_ServiceDesc, srv)
}

func _Playlist_Play_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaylistServer).Play(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Playlist/Play",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaylistServer).Play(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Playlist_Pause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaylistServer).Pause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Playlist/Pause",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaylistServer).Pause(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Playlist_Next_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaylistServer).Next(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Playlist/Next",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaylistServer).Next(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Playlist_Prev_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaylistServer).Prev(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Playlist/Prev",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaylistServer).Prev(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Playlist_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSongRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaylistServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Playlist/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaylistServer).Create(ctx, req.(*CreateSongRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Playlist_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSongRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaylistServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Playlist/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaylistServer).Delete(ctx, req.(*DeleteSongRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Playlist_ReadSong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaylistServer).ReadSong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Playlist/ReadSong",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaylistServer).ReadSong(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Playlist_ReadPlaylist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaylistServer).ReadPlaylist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Playlist/ReadPlaylist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaylistServer).ReadPlaylist(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Playlist_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSongRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaylistServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Playlist/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaylistServer).Update(ctx, req.(*UpdateSongRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Playlist_ServiceDesc is the grpc.ServiceDesc for Playlist service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Playlist_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Playlist",
	HandlerType: (*PlaylistServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Play",
			Handler:    _Playlist_Play_Handler,
		},
		{
			MethodName: "Pause",
			Handler:    _Playlist_Pause_Handler,
		},
		{
			MethodName: "Next",
			Handler:    _Playlist_Next_Handler,
		},
		{
			MethodName: "Prev",
			Handler:    _Playlist_Prev_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Playlist_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Playlist_Delete_Handler,
		},
		{
			MethodName: "ReadSong",
			Handler:    _Playlist_ReadSong_Handler,
		},
		{
			MethodName: "ReadPlaylist",
			Handler:    _Playlist_ReadPlaylist_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Playlist_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Part_Two/proto/route.proto",
}
