package graph

import (
	"strconv"
	"strings"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
)

type Graph struct {
	*simple.UndirectedGraph
}

func New() Graph {
	return Graph{simple.NewUndirectedGraph()}
}

func (g *Graph) InsertNodes(n ...int64) (success string) {
	var new bool
	var node graph.Node

	success = "Added:"

	for _, i := range n {
		node, new = g.NodeWithID(i)
		str := " " + strconv.Itoa(int(i))
		if new {
			g.AddNode(node)
			success += str
		}
	}

	return success
}

func (g *Graph) RemoveNodes(n ...int64) (success string) {
	var new bool

	success = "Removed:"

	for _, i := range n {
		_, new = g.NodeWithID(i)
		str := " " + strconv.Itoa(int(i))
		if !new {
			g.RemoveNode(i)
			success += str
		}
	}

	return success
}

func (g *Graph) NodeList() []graph.Node {
	return graph.NodesOf(g.Nodes())
}

func (g *Graph) NodeListStr() string {
	nodes := g.NodeList()
	nodesStr := make([]string, len(nodes))
	for i, n := range nodes {
		nodesStr[i] = strconv.Itoa(int(n.ID()))
	}
	return strings.Join(nodesStr, " ")
}
