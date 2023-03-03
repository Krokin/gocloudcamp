package playlist

import (
	"os"
	"reflect"
	"sync"
	"time"

	"github.com/Krokin/gocloudcamp/Part_Two/server/errors"

	jsoniter "github.com/json-iterator/go"
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
	ReadPlaylist() ([]Song, error)
	ReadSong() (*Song, error)
	UpdateInfo(title string, author string, data Song) error
	DeleteSong(title string, author string) error
	CreateSong(title string, author string, dur time.Duration) error
	SavePlaylist(path string) error
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

func NewPlaylist() *SongsPlaylist {
	s := &SongsPlaylist{control: make(chan int, 10)}
	go ControlPlay(s)
	return s
}

// эмулирует воспроизведение песни
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

type Song struct {
	Author   string        `json:"author"`
	Title    string        `json:"title"`
	Duration time.Duration `json:"duration"`
}

func NewSong(title string, author string, dur time.Duration) (*Song, error) {
	if title == "" || author == "" || dur < time.Second {
		return nil, errors.ErrorNotValid
	}
	return &Song{author, title, dur}, nil
}

func (d *SongsPlaylist) Empty() bool {
	return d.len == 0
}

func (d *SongsPlaylist) ReadSong() (*Song, error) {
	if d.Empty() {
		return nil, errors.ErrorEmptyPlaylist
	}
	d.mu.Lock()
	res := &Song{d.curr.data.Author, d.curr.data.Title, d.curr.data.Duration}
	d.mu.Unlock()
	return res, nil
}

func (d *SongsPlaylist) ReadPlaylist() ([]Song, error) {
	if d.Empty() {
		return nil, errors.ErrorEmptyPlaylist
	}
	res := make([]Song, d.len)
	d.mu.Lock()
	for i, s := 0, d.head; i < d.len; i++ {
		res[i] = *s.data
		s = s.next
	}
	d.mu.Unlock()
	return res, nil
}

func (d *SongsPlaylist) CreateSong(title string, author string, dur time.Duration) error {
	s, err := NewSong(title, author, dur)
	if err != nil {
		return err
	}
	d.addSong(s)
	return nil
}

func (d *SongsPlaylist) DeleteSong(title string, author string) error {
	node := d.head
	for node != nil {
		if node.data.Title == title && node.data.Author == author {
			break
		}
		node = node.next
	}
	if node == nil {
		return errors.ErrorNotFound
	}
	if reflect.DeepEqual(node.data, d.curr.data) && d.play {
		return errors.ErrorPlaySongNow
	}
	d.mu.Lock()
	if node == d.head && node.next != nil {
		d.head = node.next
		d.head.prev = nil
	} else if node == d.tail && node.prev != nil {
		d.tail = node.prev
		d.tail.next = nil
	} else if node.next != nil && node.prev != nil {
		node.next.prev = node.prev
		node.prev.next = node.next
	}
	d.len--
	if d.Empty() {
		d.tail = nil
		d.head = nil
	}
	d.mu.Unlock()
	if node == d.curr {
		d.Next()
	}
	return nil
}

func (d *SongsPlaylist) UpdateInfo(title string, author string, data Song) error {
	node := d.head
	for node != nil {
		if node.data.Title == title && node.data.Author == author {
			break
		}
		node = node.next
	}
	if node == nil {
		return errors.ErrorNotFound
	}
	d.mu.Lock()
	node.data = &data
	d.mu.Unlock()
	return nil
}

func (d *SongsPlaylist) Play() error {
	if d.Empty() {
		return errors.ErrorEmptyPlaylist
	}
	d.mu.Lock()
	d.play = true
	d.mu.Unlock()
	d.control <- play
	return nil
}

func (d *SongsPlaylist) Pause() error {
	if d.Empty() {
		return errors.ErrorEmptyPlaylist
	}
	d.mu.Lock()
	d.play = false
	d.mu.Unlock()
	d.control <- pause
	return nil
}

func (d *SongsPlaylist) addSong(s *Song) {
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
}

func (d *SongsPlaylist) Next() error {
	if d.Empty() {
		return errors.ErrorEmptyPlaylist
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
		return errors.ErrorEmptyPlaylist
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

func (d *SongsPlaylist) LoadPlaylist(path string) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	var songs []Song
	err = jsoniter.Unmarshal(b, &songs)
	if err != nil {
		return err
	}
	if len(songs) == 0 {
		return errors.ErrorEmptyPlaylist
	}
	for _, s := range songs {
		err = d.CreateSong(s.Title, s.Author, s.Duration)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *SongsPlaylist) SavePlaylist(path string) error {
	playlist, _ := d.ReadPlaylist()
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	res, err := jsoniter.Marshal(playlist)
	if err != nil {
		return err
	}
	_, err = f.Write(res)
	if err != nil {
		return err
	}
	return nil
}
