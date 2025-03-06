package dom

type Node struct {
	Childern  []Node
	Node_Type NodeType
}

type NodeType struct {
	Text    string
	Element ElementData
}

type ElementData struct {
	TagName string
	Attrs   AttrMap
}

type AttrMap = map[string]string
