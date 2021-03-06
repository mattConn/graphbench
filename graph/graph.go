package graph

import (
	"fmt"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/graphs/gen"
	"gonum.org/v1/gonum/graph/simple"
)

type Graph struct {
	*simple.UndirectedGraph
}

type IDList []int64

func (ids IDList) Len() int {
	return len(ids)
}

func (ids IDList) ID(i int) int64 {
	return ids[i]
}

func New() Graph {
	return Graph{simple.NewUndirectedGraph()}
}

func (g *Graph) InsertNodes(n ...int64) int {
	var new bool
	var node graph.Node

	count := 0

	for _, i := range n {
		node, new = g.NodeWithID(i)
		if new {
			g.AddNode(node)
			count++
		}
	}

	return count
}

func (g *Graph) RemoveNodes(n ...int64) int {
	count := len(g.NodeList())

	for _, i := range n {
		g.RemoveNode(i)
	}
	return count - len(g.NodeList())
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

func (g *Graph) InsertEdges(n ...int64) int {
	var from, to graph.Node

	count := 0

	for i := 0; i < len(n); i += 2 {
		if g.EdgeBetween(int64(n[i]), int64(n[i+1])) == nil {
			from, _ = g.NodeWithID(int64(n[i]))
			to, _ = g.NodeWithID(int64(n[i+1]))
			g.SetEdge(g.NewEdge(from, to))
			count++
		}
	}

	return count
}

func (g *Graph) RemoveEdges(n ...int64) int {

	count := 0

	for i := 0; i < len(n); i += 2 {
		if g.EdgeBetween(int64(n[i]), int64(n[i+1])) != nil {
			g.RemoveEdge(n[i], n[i+1])
			count++
		}
	}

	return count
}

func (g *Graph) EdgeList() []graph.Edge {
	return graph.EdgesOf(g.Edges())
}

func (g *Graph) EdgeListStr() string {
	edges := g.EdgeList()
	if len(edges) == 0 {
		return "No edges"
	}
	edgesStr := make([]string, len(edges))
	for i, e := range edges {
		edgesStr[i] = fmt.Sprintf("(%d %d)", e.From().ID(), e.To().ID())
	}
	return strings.Join(edgesStr, " ")
}

func (g *Graph) InsertCycle(n ...int64) int {
	count := len(g.EdgeList())

	gen.Cycle(g, IDList(n))

	return len(g.EdgeList()) - count
}
