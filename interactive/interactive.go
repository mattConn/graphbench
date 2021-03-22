package interactive

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type JumpTable map[string]func(*Session)

type Session struct {
	Reader *bufio.Reader
	Writer *bufio.Writer
	Input  string
	Active bool
	Cursor string
}

func NewSession() Session {
	s := Session{
		Reader: bufio.NewReader(os.Stdin),
		Writer: bufio.NewWriter(os.Stdout),
		Input:  "",
		Active: true,
		Cursor: "> ",
	}

	return s
}

func (s *Session) ReadInput() {
	input, _ := s.Reader.ReadString('\n')
	s.Input = strings.TrimSuffix(input, "\n")
}

func (s Session) WriteOutput(output string) {
	s.Writer.WriteString(output)
	s.Writer.Flush()
}

func (s *Session) ExecCommand(command string) {
	cmd, ok := Commands[command]
	if !ok {
		s.WriteOutput(fmt.Sprintf("Unrecognized command: %s\n", command))
		return
	}
	cmd(s)
}

func (s Session) HandleInput() {
	s.ReadInput()
}

var Commands = JumpTable{
	"exit": func(s *Session) {
		s.Active = false
	},
}