type Node struct {
    Children [26]*Node
    IsWordEnd bool
}

type Trie struct {
    root *Node
}


func Constructor() Trie {
    return Trie{root: &Node{}}
}

func (this *Trie) getKey(val rune) rune {
    return val - 'a'
}


func (this *Trie) Insert(word string)  {
    node := this.root
    for _, val := range word {
        key := this.getKey(val)
        if node.Children[key] == nil {
            node.Children[key] = &Node{}
        }
        node = node.Children[key]
    }
    node.IsWordEnd = true
}


func (this *Trie) Search(word string) bool {
    node := this.getLastNode(word)
    if node == nil || !node.IsWordEnd {
        return false
    }
    return true
}

func (this *Trie) getLastNode(word string) *Node {
    node := this.root
    for _, val := range word {
        key := this.getKey(val)
        if node.Children[key] == nil {
            return nil
        } 
        node = node.Children[key]
    }
    return node
}


func (this *Trie) StartsWith(prefix string) bool {
    node := this.getLastNode(prefix)
    if node == nil {
        return false
    }
    return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */