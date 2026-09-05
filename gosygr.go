package gosygr 

type Node struct {
	ValueIsSet bool
	Value string
	Deduplicated bool
	Children []*Node 
}

func (N Node) SetValue(v string) error {
	if N.ValueIsSet == false {
		N.Value = v
		N.ValueIsSet = true
		return nil
	}
	return fmt.Errorf("Value has already been set.")
}

func (N Node) UnsetValue() {
	N.ValueIsSet = false
	N.Value = ""
}

func (N Node) AppendChild(np *Node) {
	N.Children = append(N.Children, np)
}

func (N Node) InsertChild(np *Node, position int) {
	if position > len(N.Children) {
		return fmt.Errorf("Nonexisting position")
	}
	var newChildren = make([]*Node, len(N.Children) + 1)
	writen := copy(newChildren, N.Children[:position])
	if writen != position {
		return fmt.Errorf("Have copied: %d elements, expected to copy %d", writen, position)
	}
	newChildren[position] = np
	N.Children = append(newChildren, N.Children[position:])
	return nil
}



func (N Node) Deduplicate() {}

func (N Node) Reduplicate() {}


