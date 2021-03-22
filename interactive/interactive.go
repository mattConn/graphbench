package interactive

import (
	"bufio"
	"os"
)

type Session struct {
	Reader *bufio.Reader
	Writer *bufio.Writer
	Input  string
	Active bool
}

func NewSession() Session {
	s := Session{
		Reader: bufio.NewReader(os.Stdin),
		Writer: bufio.NewWriter(os.Stdout),
		Input:  "",
		Active: true,
	}

	return s
}

func (s Session) GetInput() {
	s.Writer.WriteString("> ")
	s.Writer.Flush()
	input, _ := s.Reader.ReadString('\n')
	s.Input = input
}
