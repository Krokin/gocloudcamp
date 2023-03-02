package Part_one

import (
	"reflect"
	"sync"
	"testing"
	"time"
)

func TestControlPlay(t *testing.T) {
	type args struct {
		s *SongsPlaylist
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ControlPlay(tt.args.s)
		})
	}
}

func TestNewPlaylist(t *testing.T) {
	tests := []struct {
		name string
		want *SongsPlaylist
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPlaylist(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPlaylist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSong(t *testing.T) {
	type args struct {
		title  string
		author string
		dur    time.Duration
	}
	tests := []struct {
		name    string
		args    args
		want    *Song
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewSong(tt.args.title, tt.args.author, tt.args.dur)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSong() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSong() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSongsPlaylist_AddSong(t *testing.T) {
	type fields struct {
		mu      sync.Mutex
		head    *Node
		curr    *Node
		tail    *Node
		play    bool
		control chan int
		len     int
	}
	type args struct {
		title  string
		author string
		dur    time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &SongsPlaylist{
				mu:      tt.fields.mu,
				head:    tt.fields.head,
				curr:    tt.fields.curr,
				tail:    tt.fields.tail,
				play:    tt.fields.play,
				control: tt.fields.control,
				len:     tt.fields.len,
			}
			if err := d.AddSong(tt.args.title, tt.args.author, tt.args.dur); (err != nil) != tt.wantErr {
				t.Errorf("AddSong() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSongsPlaylist_Empty(t *testing.T) {
	type fields struct {
		mu      sync.Mutex
		head    *Node
		curr    *Node
		tail    *Node
		play    bool
		control chan int
		len     int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &SongsPlaylist{
				mu:      tt.fields.mu,
				head:    tt.fields.head,
				curr:    tt.fields.curr,
				tail:    tt.fields.tail,
				play:    tt.fields.play,
				control: tt.fields.control,
				len:     tt.fields.len,
			}
			if got := d.Empty(); got != tt.want {
				t.Errorf("Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSongsPlaylist_Next(t *testing.T) {
	type fields struct {
		mu      sync.Mutex
		head    *Node
		curr    *Node
		tail    *Node
		play    bool
		control chan int
		len     int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &SongsPlaylist{
				mu:      tt.fields.mu,
				head:    tt.fields.head,
				curr:    tt.fields.curr,
				tail:    tt.fields.tail,
				play:    tt.fields.play,
				control: tt.fields.control,
				len:     tt.fields.len,
			}
			if err := d.Next(); (err != nil) != tt.wantErr {
				t.Errorf("Next() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSongsPlaylist_Pause(t *testing.T) {
	type fields struct {
		mu      sync.Mutex
		head    *Node
		curr    *Node
		tail    *Node
		play    bool
		control chan int
		len     int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &SongsPlaylist{
				mu:      tt.fields.mu,
				head:    tt.fields.head,
				curr:    tt.fields.curr,
				tail:    tt.fields.tail,
				play:    tt.fields.play,
				control: tt.fields.control,
				len:     tt.fields.len,
			}
			if err := d.Pause(); (err != nil) != tt.wantErr {
				t.Errorf("Pause() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSongsPlaylist_Play(t *testing.T) {
	type fields struct {
		mu      sync.Mutex
		head    *Node
		curr    *Node
		tail    *Node
		play    bool
		control chan int
		len     int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &SongsPlaylist{
				mu:      tt.fields.mu,
				head:    tt.fields.head,
				curr:    tt.fields.curr,
				tail:    tt.fields.tail,
				play:    tt.fields.play,
				control: tt.fields.control,
				len:     tt.fields.len,
			}
			if err := d.Play(); (err != nil) != tt.wantErr {
				t.Errorf("Play() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSongsPlaylist_Prev(t *testing.T) {
	type fields struct {
		mu      sync.Mutex
		head    *Node
		curr    *Node
		tail    *Node
		play    bool
		control chan int
		len     int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &SongsPlaylist{
				mu:      tt.fields.mu,
				head:    tt.fields.head,
				curr:    tt.fields.curr,
				tail:    tt.fields.tail,
				play:    tt.fields.play,
				control: tt.fields.control,
				len:     tt.fields.len,
			}
			if err := d.Prev(); (err != nil) != tt.wantErr {
				t.Errorf("Prev() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
