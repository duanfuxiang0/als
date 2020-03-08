package patricia

import "bytes"

type Trie struct {
	edge     []byte
	value    uint64
	depth    int
	children []*Trie
}

func LCS(key, edge []byte) (prefix, k1, k2 []byte) {
	klen, elen := len(key), len(edge)
	i := 0
	for i < klen && i < elen && key[i] == edge[i] {
		i += 1
	}
	return key[:i], key[i:], edge[i:]
}

func (t *Trie) Add(key []byte, val uint64) {
	node := t
	for {
		match := false
		for idx, tr := range node.children {
			if bytes.Equal(tr.edge, key) {
				node.children[idx].value = val
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
					node.children[idx].edge = prefix
					node.children[idx].value = val
					node.children[idx].children = []*Trie{ntr}
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
					node.children[idx].edge = prefix
					node.children[idx].value = 0
					node.children[idx].children = []*Trie{ntr1, ntr2}
					return
				}
			}
		}
		if !match {
			node.children = append(node.children, &Trie{edge: key, value: val})
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
