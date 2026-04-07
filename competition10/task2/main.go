package main

import (
	"bufio"
	"fmt"
	"os"
)

type Window struct {
	text string
}

func NewWindow() *Window {
	return &Window{text: ""}
}

func (w *Window) Add(s string) {
	w.text += s
}

func (w *Window) Backspace() {
	if len(w.text) > 0 {
		w.text = w.text[:len(w.text)-1]
	}
}

func (w *Window) Visible(k int) string {
	if len(w.text) <= k {
		return w.text
	}
	return w.text[len(w.text)-k:]
}

type Screen struct {
	windows []*Window
	cur     int
	buffer  string
	k       int
}

func NewScreen(n, k int) *Screen {
	windows := make([]*Window, n)
	for i := 0; i < n; i++ {
		windows[i] = NewWindow()
	}
	return &Screen{
		windows: windows,
		cur:     0,
		buffer:  "",
		k:       k,
	}
}

func (s *Screen) Add(ch string) {
	s.windows[s.cur].Add(ch)
}

func (s *Screen) Backspace() {
	s.windows[s.cur].Backspace()
}

func (s *Screen) Copy() {
	s.buffer = s.windows[s.cur].Visible(s.k)
}

func (s *Screen) Paste() {
	s.windows[s.cur].Add(s.buffer)
}

func (s *Screen) Next() {
	s.cur++
	if s.cur == len(s.windows) {
		s.cur = 0
	}
}

func (s *Screen) Answer() string {
	res := s.windows[s.cur].Visible(s.k)
	if res == "" {
		return "Empty"
	}
	return res
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)

	screen := NewScreen(n, k)

	for i := 0; i < m; i++ {
		var cmd string
		fmt.Fscan(in, &cmd)

		switch cmd {
		case "Backspace":
			screen.Backspace()
		case "Copy":
			screen.Copy()
		case "Paste":
			screen.Paste()
		case "Next":
			screen.Next()
		default:
			screen.Add(cmd)
		}
	}

	fmt.Println(screen.Answer())
}
