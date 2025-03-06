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

func Text(data string) Node {
	return Node{
		Childern: make([]Node, 0),
		Node_Type: NodeType{
			Text: data,
		},
	}
}

func Elem(tag_name string, attrs AttrMap, childern []Node) Node {
	return Node{
		Childern: childern,
		Node_Type: NodeType{
			Element: ElementData{
				TagName: tag_name,
				Attrs:   attrs,
			},
		},
	}
}
