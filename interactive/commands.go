package interactive

import (
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
	// "E": RemoveEdges,
}

func InsertNodes(s *Session) {
	if len(s.Operands) == 0 {
		s.WriteLine(s.GetNodeStr())
		return
	}

	s.Cache.Written = true
	IDs := make([]int64, len(s.Operands))
	for i, op := range s.Operands {
		n, err := strconv.Atoi(op)
		if err == nil {
			IDs[i] = int64(n)
		}
	}
	success := s.Graph.InsertNodes(IDs...)
	s.WriteLine(success)
}

func RemoveNodes(s *Session) {
	if len(s.Operands) == 0 {
		return
	}

	s.Cache.Written = true
	IDs := make([]int64, len(s.Operands))
	for i, op := range s.Operands {
		n, err := strconv.Atoi(op)
		if err == nil {
			IDs[i] = int64(n)
		}
	}
	success := s.Graph.RemoveNodes(IDs...)
	s.WriteLine(success)
}

func InsertEdges(s *Session) {
	if len(s.Operands) == 0 {
		s.WriteLine(s.GetEdgeStr())
		return
	}

	s.Cache.Written = true

	idLen := len(s.Operands)
	if idLen%2 != 0 {
		idLen--
	}

	IDs := make([]int64, idLen)
	for i, op := range s.Operands[:idLen] {
		n, err := strconv.Atoi(op)
		if err == nil {
			IDs[i] = int64(n)
		}
	}

	idLen = len(IDs)
	if idLen%2 != 0 {
		idLen--
	}
	success := s.Graph.InsertEdges(IDs...)
	s.WriteLine(success)
}
