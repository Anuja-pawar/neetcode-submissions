type TrieNode struct{
	children [26]*TrieNode
	end bool
}

type WordDictionary struct {
    root *TrieNode
}

func Constructor() WordDictionary {
    return WordDictionary{root: &TrieNode{}}
}

func (this *WordDictionary) AddWord(word string)  {
    curr := this.root
	for _, ch := range word{
		i := ch - 'a'
		if curr.children[i] == nil{
			curr.children[i] = &TrieNode{}
		}
		curr = curr.children[i]
	}
	curr.end = true
}

func (this *WordDictionary) Search(word string) bool {
	return this.dfs(word, 0, this.root)
}

func (this *WordDictionary) dfs(word string, j int, root *TrieNode) bool{
	curr := root
	for i := j; i < len(word); i++{
		ch := word[i]
		if ch == '.'{
			for _, child := range curr.children{
				if child != nil && this.dfs(word, i+1, child){
					return true
				}
			}
			return false
		}else{
			ind := ch - 'a'
			if curr.children[ind] == nil{
				return false
			}
			curr = curr.children[ind]
		}
	}
	return curr.end
}

