package interactive

import (
	"fmt"
	"strconv"
)

var Commands = JumpTable{
	"exit": func(s *Session) {
		s.Active = false
	},
	// node commands
	"n": InsertNodes,
	"N": RemoveNodes,
	// edge commands
	"e": InsertEdges,
	"E": RemoveEdges,
	// cycle commands
	"c": InsertCycle,
}

func InsertNodes(s *Session) {
	if len(s.Operands) == 0 {
		s.WriteLine(s.GetNodeStr())
		return
	}

	IDs := make([]int64, len(s.Operands))
	for i, op := range s.Operands {
		n, err := strconv.Atoi(op)
		if err == nil {
			IDs[i] = int64(n)
		}
	}
	inserted := s.Graph.InsertNodes(IDs...)
	if s.Verbose {
		s.WriteLine(fmt.Sprintf("+%d", inserted))
	}
}

func RemoveNodes(s *Session) {
	if len(s.Operands) == 0 {
		return
	}

	removed := 0

	if s.Operands[0] == "all" {
		for _, n := range s.Graph.NodeList() {
			s.Graph.RemoveNodes(n.ID())
			removed++
		}
		if s.Verbose {
			s.WriteLine(fmt.Sprintf("-%d", removed))
		}
		return
	}

	IDs := make([]int64, len(s.Operands))
	for i, op := range s.Operands {
		n, err := strconv.Atoi(op)
		if err == nil {
			IDs[i] = int64(n)
		}
	}
	removed = s.Graph.RemoveNodes(IDs...)
	if s.Verbose {
		s.WriteLine(fmt.Sprintf("-%d", removed))
	}
}

func InsertEdges(s *Session) {
	if len(s.Operands) == 0 {
		s.WriteLine(s.GetEdgeStr())
		return
	}

	if len(s.Operands)%2 != 0 {
		return
	}

	IDs := make([]int64, len(s.Operands))
	for i, op := range s.Operands {
		n, err := strconv.Atoi(op)
		if err == nil {
			IDs[i] = int64(n)
		}
	}

	inserted := s.Graph.InsertEdges(IDs...)
	if s.Verbose {
		s.WriteLine(fmt.Sprintf("+%d", inserted))
	}
}

func RemoveEdges(s *Session) {
	if len(s.Operands)%2 != 0 {
		return
	}

	IDs := make([]int64, len(s.Operands))
	for i, op := range s.Operands {
		n, err := strconv.Atoi(op)
		if err == nil {
			IDs[i] = int64(n)
		}
	}

	removed := s.Graph.RemoveEdges(IDs...)
	if s.Verbose {
		s.WriteLine(fmt.Sprintf("-%d", removed))
	}
}

func InsertCycle(s *Session) {
	if len(s.Operands) < 3 {
		return
	}

	IDs := make([]int64, len(s.Operands))
	for i, op := range s.Operands {
		n, err := strconv.Atoi(op)
		if err == nil {
			IDs[i] = int64(n)
		}
	}

	inserted := s.Graph.InsertCycle(IDs...)
	if s.Verbose {
		s.WriteLine(fmt.Sprintf("+%d", inserted))
	}
}
