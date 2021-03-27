package interactive

type Cache struct {
	NodeStr string
	Written bool
}

func (s *Session) GetNodeStr() string {
	c := &s.Cache
	if c.Written {
		c.Written = false
		c.NodeStr = s.Graph.NodeListStr()
	}
	return "Nodes: " + c.NodeStr
}
