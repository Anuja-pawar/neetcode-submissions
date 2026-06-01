type TrieNode struct {
	children [26]*TrieNode
	endOfWord bool
}

type PrefixTree struct {
	root *TrieNode
}

func Constructor() PrefixTree {
	return PrefixTree{root: &TrieNode{}}
}

func (this *PrefixTree) Insert(word string) {
	curr := this.root
	for _, ch := range word{
		i := ch - 'a'
		if curr.children[i] == nil{
			curr.children[i] = &TrieNode{}
		}
		curr = curr.children[i]
	}
	curr.endOfWord = true
}

func (this *PrefixTree) Search(word string) bool {
	curr := this.root
	for _, ch := range word{
		i := ch - 'a'
		if curr.children[i] == nil{
			return false
		}
		curr = curr.children[i]
	}
	return curr.endOfWord
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	curr := this.root
	for _, ch := range prefix{
		i := ch - 'a'
		if curr.children[i] == nil{
			return false
		}
		curr = curr.children[i]
	}
	return true

}
