package interactive

type Cache struct {
	NodeStr, EdgeStr     string
	NodeCount, EdgeCount int
	Written              bool
}

func NewCache() Cache {
	return Cache{
		NodeStr: "No nodes",
		EdgeStr: "No edges",
	}
}

func (s *Session) GetNodeStr() string {
	return s.Graph.NodeListStr()
}

func (s *Session) GetEdgeStr() string {
	return s.Graph.EdgeListStr()
}
