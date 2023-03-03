package Part_One

import (
	"reflect"
	"sync"
	"testing"
	"time"
)

func TestNewPlaylist(t *testing.T) {
	tests := []struct {
		name    string
		notWant *SongsPlaylist
	}{
		{
			"Test new playlist eq nil",
			nil,
		}, {
			"Test new playlist eq empty playlist",
			&SongsPlaylist{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPlaylist(); reflect.DeepEqual(got, tt.notWant) {
				t.Errorf("NewPlaylist() = %v, want %v", got, tt.notWant)
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
		{
			"Test new song valid",
			args{"A", "Boris", time.Second * 100},
			&Song{"Boris", "A", time.Second * 100},
			false,
		}, {
			"Test new song not valid author",
			args{"A", "", time.Second * 100},
			nil,
			true,
		}, {
			"Test new song not valid title",
			args{"", "Boris", time.Second * 100},
			nil,
			true,
		}, {
			"Test new song not valid duration",
			args{"A", "Boris", time.Microsecond * 100},
			nil,
			true,
		},
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
		a string
		t string
		d time.Duration
	}

	tests := []struct {
		name   string
		fields fields
		args
		wantCurr     *Node
		wantHead     *Node
		wantTailNext *Node
		wantTail     *Node
		wantCurrPrev *Node
		wantLen      int
		wantErr      bool
	}{
		{
			"Test AddSong not valid data author and title",
			fields{},
			args{a: "", t: "", d: time.Second * 100},
			nil,
			nil,
			nil,
			nil,
			nil,
			0,
			true,
		}, {
			"Test AddSong not valid data duration",
			fields{},
			args{a: "asd", t: "asd", d: time.Nanosecond},
			nil,
			nil,
			nil,
			nil,
			nil,
			0,
			true,
		}, {
			"Test AddSong empty playlist ok",
			fields{len: 0},
			args{a: "asd", t: "asd", d: time.Second * 5},
			&Node{data: &Song{"asd", "asd", time.Second * 5}},
			&Node{data: &Song{"asd", "asd", time.Second * 5}},
			nil,
			&Node{data: &Song{"asd", "asd", time.Second * 5}},
			nil,
			1,
			false,
		},
		{
			"Test AddSong not empty playlist ok",
			fields{
				len:  1,
				head: &Node{data: &Song{"t", "s", time.Second * 10}},
				curr: &Node{data: &Song{"t", "s", time.Second * 10}},
				tail: &Node{data: &Song{"t", "s", time.Second * 10}},
			},
			args{a: "asd", t: "asd", d: time.Second * 5},
			&Node{data: &Song{"t", "s", time.Second * 10}},
			&Node{data: &Song{"t", "s", time.Second * 10}},
			&Node{data: &Song{"asd", "asd", time.Second * 5}},
			&Node{data: &Song{"asd", "asd", time.Second * 5}},
			&Node{data: &Song{"t", "s", time.Second * 10}},
			2,
			false,
		},
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
			lastNode := d.tail
			if err := d.AddSong(tt.t, tt.a, tt.d); (err != nil) != tt.wantErr {
				t.Errorf("Next() error = %v, wantErr %v", err, tt.wantErr)
			} else if err != nil {
				return
			}
			if d.len != tt.wantLen {
				t.Errorf("Next() error len")
			}
			if !reflect.DeepEqual(d.curr, tt.wantCurr) {
				t.Errorf("Next() error curr")
			}
			if !reflect.DeepEqual(d.head, tt.wantHead) {
				t.Errorf("Next() error head")
			}
			if !reflect.DeepEqual(d.tail.data, tt.wantTail.data) {
				t.Errorf("Next() error tail")
			}
			if d.tail.prev != nil {
				if !reflect.DeepEqual(d.tail.prev.data, tt.wantCurrPrev.data) {
					t.Errorf("Next() error curr prev")
				}
			} else if d.tail.prev == nil && tt.wantCurrPrev != nil {
				t.Errorf("Next() error curr prev")
			}
			if lastNode != nil {
				if !reflect.DeepEqual(d.tail.prev.next.data, tt.wantTailNext.data) {
					t.Errorf("Next() error tail next")
				}
			}
		})
	}
}

func TestSongsPlaylist_Empty(t *testing.T) {
	tests := []struct {
		name   string
		fields *SongsPlaylist
		want   bool
	}{
		{
			"Test empty true",
			&SongsPlaylist{},
			true,
		}, {
			"Test empty false",
			&SongsPlaylist{len: 1},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.Empty(); got != tt.want {
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
		name     string
		fields   fields
		wantCurr *Node
		wantErr  bool
	}{
		{
			"Test next empty playlist",
			fields{len: 0, control: make(chan int, 1)},
			nil,
			true,
		}, {
			"Test next nil next",
			fields{head: &Node{data: &Song{Author: "ok"}}, curr: &Node{}, len: 2, control: make(chan int, 1)},
			&Node{data: &Song{Author: "ok"}},
			false,
		}, {
			"Test next next curr ok",
			fields{curr: &Node{next: &Node{data: &Song{Author: "ok"}}}, len: 2, control: make(chan int, 1)},
			&Node{data: &Song{Author: "ok"}},
			false,
		},
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
			} else if err != nil {
				return
			}
			if d.curr.data.Author != tt.wantCurr.data.Author {
				t.Errorf("Next() error curr field")
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
		name        string
		fields      fields
		wantPlay    bool
		wantControl int
		wantErr     bool
	}{
		{
			"Test pause ok",
			fields{play: true, len: 1, control: make(chan int, 1)},
			false,
			2,
			false,
		}, {
			"Test pause empty playlist",
			fields{control: make(chan int, 1)},
			false,
			0,
			true,
		},
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
			err := d.Pause()
			if (err != nil) != tt.wantErr {
				t.Errorf("Pause() error = %v, wantErr %v", err, tt.wantErr)
			}
			if d.play != tt.wantPlay {
				t.Errorf("Pause() error field play")
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
		name     string
		fields   fields
		wantPlay bool
		wantErr  bool
	}{
		{
			"Test play ok",
			fields{curr: &Node{}, play: false, len: 1, control: make(chan int, 1)},
			true,
			false,
		}, {
			"Test play empty playlist",
			fields{len: 0, control: make(chan int, 1)},
			false,
			true,
		},
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
			} else if err != nil {
				return
			}
			if d.play != tt.wantPlay {
				t.Errorf("Play() error field play")
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
		name     string
		fields   fields
		wantCurr *Node
		wantErr  bool
	}{
		{
			"Test prev empty playlist",
			fields{len: 0, control: make(chan int, 1)},
			nil,
			true,
		}, {
			"Test prev nil prev",
			fields{tail: &Node{data: &Song{Author: "ok"}}, curr: &Node{}, len: 2, control: make(chan int, 1)},
			&Node{data: &Song{Author: "ok"}},
			false,
		}, {
			"Test prev prev curr ok",
			fields{curr: &Node{prev: &Node{data: &Song{Author: "ok"}}}, len: 2, control: make(chan int, 1)},
			&Node{data: &Song{Author: "ok"}},
			false,
		},
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
			} else if err != nil {
				return
			}
			if d.curr.data.Author != tt.wantCurr.data.Author {
				t.Errorf("Prev() error curr field")
			}
		})
	}
}
