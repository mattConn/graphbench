package interactive

import (
	"bufio"
	"fmt"
	"graphbench/graph"
	"os"
	"strings"
)

type JumpTable map[string]func(*Session)

type Session struct {
	Reader   *bufio.Reader
	Writer   *bufio.Writer
	Input    string
	Operands []string
	Active   bool
	Cursor   string
	Graph    graph.Graph
	Cache    Cache
}

func NewSession() Session {
	s := Session{
		Reader: bufio.NewReader(os.Stdin),
		Writer: bufio.NewWriter(os.Stdout),
		Input:  "",
		Active: true,
		Cursor: "> ",
		Graph:  graph.New(),
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

func (s Session) WriteLine(output string) {
	s.WriteOutput(output + "\n")
}

func (s *Session) ExecCommand(commandString string) {
	if len(commandString) == 0 {
		return
	}
	commands := strings.Fields(commandString)
	cmd, ok := Commands[commands[0]]
	if !ok {
		s.WriteOutput(fmt.Sprintf("Unrecognized command: %s\n", commands[0]))
		return
	}
	s.Operands = commands[1:]
	cmd(s)
}

func (s Session) Run() {
	for s.Active {
		s.WriteOutput(s.Cursor)
		s.ReadInput()
		s.ExecCommand(s.Input)
	}
}
