package main

import (
	"graphbench/interactive"
	"os"
)

func main() {
	s := interactive.NewSession()
	for _, arg := range os.Args[1:] {
		switch arg {
		case "-v":
			s.Verbose = true
		}
	}
	s.Run()
}
