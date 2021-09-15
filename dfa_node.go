package sensitive_words_filter

type DFANode struct {
	End  bool
	Next map[rune]*DFANode
}

func (n *DFANode) AddWord(word string) {
	node := n
	chars := []rune(word)
	for index, _ := range chars {
		node = node.AddChild(chars[index])
	}
	node.End = true
}

func (n *DFANode) AddChild(c rune) *DFANode {
	if n.Next == nil {
		n.Next = make(map[rune]*DFANode)
	}

	//如果已经存在了，就不再往里面添加了；
	if next, ok := n.Next[c]; ok {
		return next
	} else {
		n.Next[c] = &DFANode{
			End:  false,
			Next: nil,
		}
		return n.Next[c]
	}
}

func (n *DFANode) FindChild(c rune) *DFANode {
	if n.Next == nil {
		return nil
	}

	if _, ok := n.Next[c]; ok {
		return n.Next[c]
	}
	return nil
}
