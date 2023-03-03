package Part_One

import (
	"sync"
	"time"
)

const (
	play = iota + 1
	pause
	next
	prev
	TickPlay = time.Microsecond * 100000
)

type Playlister interface {
	Play() error
	Pause() error
	Next() error
	Prev() error
	AddSong(title string, author string, dur time.Duration) error
}

type SongsPlaylist struct {
	mu      sync.Mutex
	head    *Node
	curr    *Node
	tail    *Node
	play    bool
	control chan int
	len     int
}

type Node struct {
	next *Node
	prev *Node
	data *Song
}

type Song struct {
	Author   string
	Title    string
	Duration time.Duration
}

func NewPlaylist() *SongsPlaylist {
	s := &SongsPlaylist{control: make(chan int, 10)}
	go ControlPlay(s)
	return s
}

func NewSong(title string, author string, dur time.Duration) (*Song, error) {
	if title == "" || author == "" || dur < time.Second {
		return nil, ErrorNotValid
	}
	return &Song{author, title, dur}, nil
}

func ControlPlay(s *SongsPlaylist) {
	var t = new(time.Ticker)
	var dur time.Duration
	for {
		select {
		case <-t.C:
			dur -= TickPlay
			if dur <= 0 {
				s.Next()
			}
		case c := <-s.control:
			switch c {
			case play:
				if dur == 0 {
					s.mu.Lock()
					dur = s.curr.data.Duration
					s.mu.Unlock()
				}
				t = time.NewTicker(TickPlay)
			case pause:
				t.Stop()
			case next, prev:
				s.mu.Lock()
				dur = s.curr.data.Duration
				s.mu.Unlock()
			}
		}
	}
}

func (d *SongsPlaylist) Play() error {
	if d.Empty() {
		return ErrorEmptyPlaylist
	}
	d.mu.Lock()
	d.play = true
	d.mu.Unlock()
	d.control <- play
	return nil
}

func (d *SongsPlaylist) Pause() error {
	if d.Empty() {
		return ErrorEmptyPlaylist
	}
	d.mu.Lock()
	d.play = false
	d.mu.Unlock()
	d.control <- pause
	return nil
}

func (d *SongsPlaylist) AddSong(title string, author string, dur time.Duration) error {
	s, err := NewSong(title, author, dur)
	if err != nil {
		return err
	}
	newNode := &Node{data: s}
	d.mu.Lock()
	if d.tail == nil {
		d.head = newNode
		d.curr = newNode
	} else {
		d.tail.next = newNode
		newNode.prev = d.tail
	}
	d.tail = newNode
	d.len++
	d.mu.Unlock()
	return nil
}

func (d *SongsPlaylist) Next() error {
	if d.Empty() {
		return ErrorEmptyPlaylist
	}
	d.mu.Lock()
	if d.curr.next == nil {
		d.curr = d.head
	} else {
		d.curr = d.curr.next
	}
	d.mu.Unlock()
	d.control <- next
	return nil
}

func (d *SongsPlaylist) Prev() error {
	if d.Empty() {
		return ErrorEmptyPlaylist
	}
	d.mu.Lock()
	if d.curr.prev == nil {
		d.curr = d.tail
	} else {
		d.curr = d.curr.prev
	}
	d.mu.Unlock()
	d.control <- prev
	return nil
}

func (d *SongsPlaylist) Empty() bool {
	return d.len == 0
}
