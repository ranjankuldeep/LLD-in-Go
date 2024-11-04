package trie

import (
	"container/list"
	"sync"
)

type Bytes []byte

type trieNode struct {
	children map[byte]*trieNode
	flag     bool
	symbol   byte
	root     bool
	value    []byte
}

// Trie Thread safe implementation of "Trie" data structure
type Trie struct {
	root *trieNode
	rw   sync.RWMutex
	size int
}

func NewTrie() *Trie {
	return &Trie{
		root: &trieNode{root: true, children: make(map[byte]*trieNode)},
		size: 1,
	}
}

func newNode(symbol byte) *trieNode {
	return &trieNode{children: make(map[byte]*trieNode),
		symbol: symbol}
}

func (t *Trie) Size() int {
	t.rw.RLock()
	defer t.rw.RUnlock()
	return t.size
}

func (t *Trie) insert(key, val Bytes) {
	t.rw.Lock()
	defer t.rw.Unlock()

	currNode := t.root
	for _, symbol := range key {
		if currNode.children[symbol] == nil {
			currNode.children[symbol] = newNode(symbol)
		}
		currNode = currNode.children[symbol]
	}

	if currNode.value == nil {
		t.size++
	}
	currNode.value = val
}

func (t *Trie) search(key []byte) ([]byte, bool) {
	t.rw.RLock()
	defer t.rw.RUnlock()

	currNode := t.root

	for _, symbol := range key {
		if currNode.children[symbol] == nil {
			return nil, false
		}
		currNode = currNode.children[symbol]
	}
	return currNode.value, true
}

func (t *Trie) GetAllKeys() []Bytes {
	t.rw.RLock()
	defer t.rw.RUnlock()

	visited := make(map[*trieNode]bool)
	keys := make([]Bytes, 0, t.size)

	var dfsGetKeys func(n *trieNode, key Bytes)
	dfsGetKeys = func(n *trieNode, key Bytes) {
		if n != nil {
			pathKey := append(key, n.symbol)
			visited[n] = true

			if n.value != nil {
				fullKey := make(Bytes, len(pathKey))
				copy(fullKey, pathKey)
				keys = append(keys, fullKey[1:])
			}

			for _, child := range n.children {
				if _, ok := visited[child]; !ok {
					dfsGetKeys(child, pathKey)
				}
			}
		}
	}
	dfsGetKeys(t.root, Bytes{})
	return keys
}

func (t *Trie) GetAllValues() []Bytes {
	t.rw.RLock()
	defer t.rw.RUnlock()
	queue := list.New()
	values := make([]Bytes, 0, t.size)
	visited := make(map[*trieNode]bool)

	queue.PushBack(t.root)
	for queue.Len() > 0 {
		element := queue.Front()
		queue.Remove(element)

		node := element.Value.(*trieNode)
		visited[node] = true

		for _, child := range node.children {
			_, ok := visited[child]
			if !ok {
				queue.PushBack(child)
			}
		}

		if node.value != nil {
			values = append(values, node.value)
		}
	}
	return values
}

func (t *Trie) GetPrefixKeys(prefix Bytes) []Bytes {
	t.rw.RLock()
	defer t.rw.RUnlock()

	keys := make([]Bytes, 0, t.size)
	visited := make(map[*trieNode]bool)

	if len(prefix) == 0 {
		return keys
	}

	var dfsGetPrefixKeys func(n *trieNode, prefixIdx int, key Bytes)
	dfsGetPrefixKeys = func(n *trieNode, prefixIdx int, key Bytes) {
		if n != nil {
			pathKey := append(key, n.symbol)
			if prefixIdx == len(prefix) || n.symbol == prefix[prefixIdx] {
				visited[n] = true

				if n.value != nil {
					fullKey := make([]byte, len(pathKey))
					copy(fullKey, pathKey)
					keys = append(keys, fullKey)
				}

				if prefixIdx < len(prefix) {
					prefixIdx++
				}

				for _, child := range n.children {
					if _, ok := visited[child]; !ok {
						dfsGetPrefixKeys(child, prefixIdx, pathKey)
					}
				}
			}
		}
	}
	if n, ok := t.root.children[prefix[0]]; ok {
		dfsGetPrefixKeys(n, 0, Bytes{})
	}
	return keys
}

func (t *Trie) GetPrefixValues(prefix Bytes) []Bytes {
	t.rw.RLock()
	defer t.rw.RUnlock()

	values := make([]Bytes, 0, t.size)
	visited := make(map[*trieNode]bool)

	if len(prefix) == 0 {
		return values
	}

	var dfsGetPrefixValues func(n *trieNode, prefixIdx int)
	dfsGetPrefixValues = func(n *trieNode, prefixIdx int) {
		if n != nil {
			if prefixIdx == len(prefix) || n.symbol == prefix[prefixIdx] {
				visited[n] = true
				if n.value != nil {
					values = append(values, n.value)
				}

				if prefixIdx < len(prefix) {
					prefixIdx++
				}

				for _, child := range n.children {
					if _, ok := visited[child]; !ok {
						dfsGetPrefixValues(child, prefixIdx)
					}
				}
			}
		}
	}
	if n, ok := t.root.children[prefix[0]]; ok {
		dfsGetPrefixValues(n, 0)
	}
	return values
}
