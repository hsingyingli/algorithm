package trie

type Node struct {
	dict map[byte]*Node
	end  bool
}

func (node *Node) Insert(str string) {
	character := []byte(str)

	current := node
	for _, c := range character {
		next, ok := current.dict[c]
		if !ok {
			next = &Node{
				dict: make(map[byte]*Node),
				end:  false,
			}
			current.dict[c] = next
		}

		current = next
	}

	current.end = true
}

func (node *Node) Find(str string) bool {
	character := []byte(str)
	current := node

	for _, c := range character {
		next, ok := current.dict[c]

		if !ok {
			return false
		}
		current = next
	}
	return current.end
}
