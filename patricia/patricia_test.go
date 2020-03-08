package patricia

import (
	"fmt"
	"testing"
)

func TestLCP(t *testing.T) {
	e := []byte("another")
	k := []byte("anwaytogo")
	p, k1, k2 := LCS(e, k)
	fmt.Println(string(p), string(k1), string(k2))
}

func TestTrie_Add(t *testing.T) {
	trie := &Trie{}
	trie.Add([]byte("a"), 1)
	trie.Add([]byte("an"), 2)
	trie.Add([]byte("another"), 3)
	trie.Add([]byte("bool"), 4)
	trie.Add([]byte("boy"), 5)
	trie.Add([]byte("zoo"), 6)
	fmt.Printf("%v", *trie)
}

func TestTrie_Find(t *testing.T) {
	trie := &Trie{}
	trie.Add([]byte("a"), 1)
	trie.Add([]byte("an"), 2)
	trie.Add([]byte("another"), 3)
	trie.Add([]byte("bool"), 4)
	trie.Add([]byte("boy"), 5)
	trie.Add([]byte("zoo"), 6)

	fmt.Println(trie.Find([]byte("zoo1")))
}