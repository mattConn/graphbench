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
	c := &s.Cache
	if c.Written {
		c.Written = false
		c.NodeStr = s.Graph.NodeListStr()
	}
	return c.NodeStr
}

func (s *Session) GetEdgeStr() string {
	c := &s.Cache
	if c.Written {
		c.Written = false
		c.EdgeStr = s.Graph.EdgeListStr()
	}
	return c.EdgeStr
}
