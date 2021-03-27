package graph

import (
	"fmt"
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

	success = "Added"

	for _, i := range n {
		node, new = g.NodeWithID(i)
		str := fmt.Sprintf(" %s", strconv.Itoa(int(i)))
		if new {
			g.AddNode(node)
			success += str
		}
	}

	return success
}

func (g *Graph) RemoveNodes(n ...int64) (success string) {
	var new bool

	success = "Removed"

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
	if len(nodes) == 0 {
		return "No nodes"
	}
	nodesStr := make([]string, len(nodes))
	for i, n := range nodes {
		nodesStr[i] = strconv.Itoa(int(n.ID()))
	}
	return strings.Join(nodesStr, " ")
}

func (g *Graph) InsertEdges(n ...int64) (success string) {
	var from, to graph.Node

	success = "Added"

	for i := 0; i < len(n); i += 2 {
		from, _ = g.NodeWithID(int64(n[i]))
		to, _ = g.NodeWithID(int64(n[i+1]))
		g.SetEdge(g.NewEdge(from, to))
		success += fmt.Sprintf(" (%d %d)", n[i], n[i+1])
	}

	return success
}

func (g *Graph) EdgeList() []graph.Edge {
	return graph.EdgesOf(g.Edges())
}

func (g *Graph) EdgeListStr() string {
	edges := g.EdgeList()
	if len(edges) == 0 {
		return "No nodes"
	}
	edgesStr := make([]string, len(edges))
	for i, e := range edges {
		edgesStr[i] = fmt.Sprintf("(%d %d)", e.From().ID(), e.To().ID())
	}
	return strings.Join(edgesStr, " ")
}
