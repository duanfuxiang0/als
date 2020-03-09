package patricia

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
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

func TestBigKey(t *testing.T) {
	trie := &Trie{}
	f, err := os.Open("../gen.log")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	buf := bufio.NewReader(f)
	n := 1
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF{
				fmt.Println("file end")
				break
			} else {
				fmt.Println(err)
				break
			}
		}
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		trie.Add([]byte(line), uint64(n))
		n = n + 1
	}
	fmt.Println(trie.Find([]byte("YvASKQp")))
	fmt.Println(trie.Find([]byte("CvE4YBdABwZtN9ujUWeDx1v6")))
	fmt.Println(trie.Find([]byte("MvpTLLSu1NznBDG9qM1FMUDwLLw6C8NQxKaL3xvXbNUqzxgyUQt2PuxE")))
	fmt.Println(trie.Find([]byte("sWWkPPJ9gSyoLFn4kTNyuT9Pz8ztD9Ql")))
}