package patricia

import (
	"bytes"
	"fmt"
	"sync"
)

type Trie struct {
	lock     sync.Mutex
	edge     []byte
	value    uint64
	children []*Trie
}

func LCS(key1, key2 []byte) (prefix, k1, k2 []byte) {
	len1, len2 := len(key1), len(key2)
	i := 0
	for i < len1 && i < len2 && key1[i] == key2[i] {
		i += 1
	}
	return key1[:i], key1[i:], key2[i:]
}

func (t *Trie) Add(key []byte, val uint64) {
	if len(key) == 0 {
		fmt.Printf("add empty key...\n")
		return
	}
	node := t
	for {
		match := false
		for idx, tr := range node.children {
			if bytes.Equal(tr.edge, key) {
				node.lock.Lock()
				node.children[idx].value = val
				node.lock.Unlock()
				return
			}
			prefix, k1, k2 := LCS(tr.edge, key)
			if len(prefix) != 0 {
				match = true
				if len(k1) == 0 {
					node = tr
					key = k2
					break
				} else if len(k2) == 0 {
					ntr := &Trie{
						edge:  k1,
						value: tr.value,
					}
					node.lock.Lock()
					node.children[idx].edge = prefix
					node.children[idx].value = val
					node.children[idx].children = []*Trie{ntr}
					node.lock.Unlock()
					return
				} else {
					ntr1 := &Trie{
						edge:  k1,
						value: tr.value,
					}
					ntr2 := &Trie{
						edge:  k2,
						value: val,
					}
					node.lock.Lock()
					node.children[idx].edge = prefix
					node.children[idx].value = 0
					node.children[idx].children = []*Trie{ntr1, ntr2}
					node.lock.Unlock()
					return
				}
			}
		}
		if !match {
			node.lock.Lock()
			node.children = append(node.children, &Trie{edge: key, value: val})
			node.lock.Unlock()
			return
		}
	}
}

func (t *Trie) Find(key []byte) uint64 {
	tree := t
	for {
		match := false
		for _, tr := range tree.children {
			if bytes.Equal(tr.edge, key) {
				return tr.value
			}
			prefix, k1, k2 := LCS(tr.edge, key)
			if len(prefix) != 0 && len(k1) == 0 {
				match = true
				key = k2
				tree = tr
				break
			}
		}
		if !match {
			return 0
		}
	}
}
