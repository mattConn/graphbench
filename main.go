package main

import "graph-serve/interactive"

func main() {
	s := interactive.NewSession()
	for s.Active {
		s.GetInput()
	}
}
