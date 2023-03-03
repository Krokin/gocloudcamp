package playlist

import (
	"reflect"
	"sync"
	"testing"
	"time"
)

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

func TestSongsPlaylist_CreateSong(t *testing.T) {
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
		{
			"Test CreateSong not valid author or title",
			fields{},
			args{"", "", time.Second * 10},
			true,
		}, {
			"Test CreateSong not valid duration",
			fields{},
			args{"", "", time.Nanosecond},
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
			if err := d.CreateSong(tt.args.title, tt.args.author, tt.args.dur); (err != nil) != tt.wantErr {
				t.Errorf("CreateSong() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSongsPlaylist_DeleteSong(t *testing.T) {
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
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantNext *Node
		wantPrev *Node
		wantLen  int
		wantErr  bool
	}{
		{
			"Test DeleteSong empty playlist",
			fields{},
			args{},
			nil,
			nil,
			0,
			true,
		}, {
			"Test DeleteSong not found song in playlist",
			fields{len: 2,
				head: &Node{
					data: &Song{Author: "s", Title: "a", Duration: time.Second * 10},
					next: &Node{
						data: &Song{Author: "s", Title: "b", Duration: time.Second * 15}},
				},
			},
			args{"aaa", "sss"},
			nil,
			nil,
			2,
			true,
		}, {
			"Test DeleteSong play song",
			fields{len: 2,
				head: &Node{
					data: &Song{Author: "s", Title: "a", Duration: time.Second * 10},
					next: &Node{
						data: &Song{Author: "s", Title: "b", Duration: time.Second * 15}},
				},
				curr: &Node{data: &Song{Author: "s", Title: "b", Duration: time.Second * 15}},
				play: true,
			},
			args{"b", "s"},
			nil,
			nil,
			2,
			true,
		}, {
			"Test DeleteSong last song",
			fields{len: 1,
				head: &Node{data: &Song{Author: "s", Title: "b", Duration: time.Second * 15}},
				curr: &Node{data: &Song{Author: "s", Title: "b", Duration: time.Second * 15}},
			},
			args{"b", "s"},
			nil,
			nil,
			0,
			false,
		}, {
			"Test DeleteSong in middle playlist",
			fields{len: 3,
				head: &Node{
					data: &Song{Author: "s", Title: "a", Duration: time.Second * 10},
					next: &Node{
						data: &Song{Author: "ss", Title: "a", Duration: time.Second * 15},
						next: &Node{
							data: &Song{Author: "sss", Title: "a", Duration: time.Second * 20}},
						prev: &Node{data: &Song{Author: "s", Title: "a", Duration: time.Second * 10}},
					},
				},
				curr: &Node{data: &Song{Author: "s", Title: "b", Duration: time.Second * 15}},
			},
			args{"a", "ss"},
			&Node{data: &Song{Author: "sss", Title: "a", Duration: time.Second * 20}},
			&Node{data: &Song{Author: "s", Title: "a", Duration: time.Second * 10}},
			2,
			false,
		}, {
			"Test DeleteSong tail in playlist",
			fields{len: 2,
				head: &Node{
					data: &Song{Author: "s", Title: "a", Duration: time.Second * 10},
					next: &Node{
						data: &Song{Author: "ss", Title: "a", Duration: time.Second * 15},
					},
				},
				curr: &Node{data: &Song{Author: "s", Title: "b", Duration: time.Second * 15}},
			},
			args{"a", "ss"},
			nil,
			nil,
			1,
			false,
		}, {
			"Test DeleteSong head in playlist",
			fields{len: 2,
				head: &Node{
					data: &Song{Author: "s", Title: "a", Duration: time.Second * 10},
					next: &Node{
						data: &Song{Author: "ss", Title: "a", Duration: time.Second * 15},
					},
				},
				curr: &Node{data: &Song{Author: "s", Title: "b", Duration: time.Second * 15}},
			},
			args{"a", "s"},
			nil,
			nil,
			1,
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
			prev, next := func(title string, author string) (*Node, *Node) {
				n := d.head
				for n != nil {
					if n.data.Author == author && n.data.Title == title {
						return n.prev, n.next
					}
					n = n.next
				}
				return nil, nil
			}(tt.args.title, tt.args.author)
			if err := d.DeleteSong(tt.args.title, tt.args.author); (err != nil) != tt.wantErr {
				t.Errorf("DeleteSong() error = %v, wantErr %v", err, tt.wantErr)
			}
			if d.len != tt.wantLen {
				t.Errorf("DeleteSong() error len")
			}
			if d.len == 1 {
				prev, next = nil, nil
			}
			if tt.wantNext != nil && next != nil {
				if !reflect.DeepEqual(tt.wantNext.data, next.data) {
					t.Errorf("DeleteSong() error prev next")
				}
			} else if (tt.wantNext == nil && next != nil) || (tt.wantNext != nil && next == nil) {
				t.Errorf("DeleteSong() error prev next")
			}
			if tt.wantPrev != nil && prev != nil {
				if !reflect.DeepEqual(tt.wantPrev.data, prev.data) {
					t.Errorf("DeleteSong() error next prev")
				}
			} else if (tt.wantPrev == nil && prev != nil) || (tt.wantPrev != nil && prev == nil) {
				t.Errorf("DeleteSong() error next prev")
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

func TestSongsPlaylist_LoadPlaylist(t *testing.T) {
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
		path    string
		fields  fields
		wantErr bool
	}{
		{
			"Test LoadPlaylist error file",
			"asd",
			fields{},
			true,
		}, {
			"Test LoadPlaylist error valid data",
			"playlist.go",
			fields{},
			true,
		}, {
			"Test LoadPlaylist ok",
			"playlist_test_files/normal.json",
			fields{},
			false,
		}, {
			"Test LoadPlaylist empty file",
			"playlist_test_files/empty.json",
			fields{},
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
			if err := d.LoadPlaylist(tt.path); (err != nil) != tt.wantErr {
				t.Errorf("LoadPlaylist() error = %v, wantErr %v", err, tt.wantErr)
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

func TestSongsPlaylist_ReadPlaylist(t *testing.T) {
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
		want    []Song
		wantErr bool
	}{
		{
			"Test ReadPlaylist ok",
			fields{len: 2,
				head: &Node{
					data: &Song{Author: "s", Title: "a", Duration: time.Second * 10},
					next: &Node{
						data: &Song{Author: "s", Title: "b", Duration: time.Second * 15}},
				},
			},
			[]Song{
				{Author: "s", Title: "a", Duration: time.Second * 10},
				{Author: "s", Title: "b", Duration: time.Second * 15},
			},
			false,
		}, {
			"Test ReadPlaylist empty playlist",
			fields{},
			nil,
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
			got, err := d.ReadPlaylist()
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

func TestSongsPlaylist_ReadSong(t *testing.T) {
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
		want    *Song
		wantErr bool
	}{
		{
			"Test ReadSong ok",
			fields{len: 1, curr: &Node{data: &Song{Author: "s", Title: "a", Duration: time.Second * 10}}},
			&Song{Author: "s", Title: "a", Duration: time.Second * 10},
			false,
		}, {
			"Test ReadSong empty playlist",
			fields{},
			nil,
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
			got, err := d.ReadSong()
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

func TestSongsPlaylist_SavePlaylist(t *testing.T) {
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
		path    string
		fields  fields
		wantErr bool
	}{
		{
			"Test SavePlaylist error file",
			"asd/asd",
			fields{},
			true,
		}, {
			"Test SavePlaylist ok",
			"playlist_test_files/normal.json",
			fields{len: 2,
				head: &Node{
					data: &Song{Author: "s", Title: "a", Duration: time.Second * 10},
					next: &Node{
						data: &Song{Author: "s", Title: "b", Duration: time.Second * 15}},
				},
			},
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
			if err := d.SavePlaylist(tt.path); (err != nil) != tt.wantErr {
				t.Errorf("SavePlaylist() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSongsPlaylist_UpdateInfo(t *testing.T) {
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
		data   Song
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantSong *Song
		wantErr  bool
	}{
		{
			"Test UpdateInfo ok",
			fields{len: 2,
				head: &Node{
					data: &Song{Author: "s", Title: "a", Duration: time.Second * 10},
					next: &Node{
						data: &Song{Author: "s", Title: "b", Duration: time.Second * 15}},
				},
			},
			args{
				title:  "a",
				author: "s",
				data:   Song{"sss", "aaa", time.Second * 100},
			},
			&Song{"sss", "aaa", time.Second * 100},
			false,
		}, {
			"Test UpdateInfo empty playlist",
			fields{},
			args{
				title:  "a",
				author: "s",
				data:   Song{"sss", "aaa", time.Second * 100},
			},
			nil,
			true,
		}, {
			"Test UpdateInfo not found song in playlist",
			fields{len: 2,
				head: &Node{
					data: &Song{Author: "s", Title: "a", Duration: time.Second * 10},
					next: &Node{
						data: &Song{Author: "s", Title: "b", Duration: time.Second * 15}},
				},
			},
			args{
				title:  "aaa",
				author: "sss",
				data:   Song{"sss", "aaa", time.Second * 100},
			},
			nil,
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
			if err := d.UpdateInfo(tt.args.title, tt.args.author, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("UpdateInfo() error = %v, wantErr %v", err, tt.wantErr)
			} else if err != nil {
				return
			}
			if d.head != nil {
				if !reflect.DeepEqual(*tt.wantSong, *d.head.data) {
					t.Errorf("UpdateInfo() error update")
				}
			}
		})
	}
}

func TestSongsPlaylist_addSong(t *testing.T) {
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
		name             string
		fields           fields
		wantTail         *Node
		wantTailPrev     *Node
		wantTailPrevNext *Node
		wantLen          int
		arg              *Song
	}{
		{
			"Test addSong in empty playlist",
			fields{},
			&Node{data: &Song{"a", "b", time.Second * 10}},
			nil,
			nil,
			1,
			&Song{"a", "b", time.Second * 10},
		}, {
			"Test addSong in not empty playlist",
			fields{len: 1, tail: &Node{data: &Song{Author: "s", Title: "a", Duration: time.Second * 10}}},
			&Node{data: &Song{"a", "b", time.Second * 10}},
			&Node{data: &Song{Author: "s", Title: "a", Duration: time.Second * 10}},
			&Node{data: &Song{"a", "b", time.Second * 10}},
			2,
			&Song{"a", "b", time.Second * 10},
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
			d.addSong(tt.arg)
			if d.len != tt.wantLen {
				t.Errorf("addSong() error len")
			}
			if !reflect.DeepEqual(d.tail.data, tt.wantTail.data) {
				t.Errorf("addSong() error tail")
			}
			if tt.wantTailPrev != nil {
				if !reflect.DeepEqual(d.tail.prev.data, tt.wantTailPrev.data) {
					t.Errorf("addSong() error tail prev")
				}
				if !reflect.DeepEqual(d.tail.prev.next.data, tt.wantTailPrevNext.data) {
					t.Errorf("addSong() error tail prev next")
				}
			}
		})
	}
}
