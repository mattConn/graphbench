package main

import (
	"graph-serve/interactive"
)

func main() {
	s := interactive.NewSession()
	for s.Active {
		s.WriteOutput(s.Cursor)
		s.ReadInput()
		s.ExecCommand(s.Input)
	}
}
