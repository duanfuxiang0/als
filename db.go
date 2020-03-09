package als

import (
	"bytes"
	"fmt"
	"os"
	"time"

	"als/patricia"
)

type DB struct {
	tr *patricia.Trie
	af *AFile
}

func (db *DB) Init(file string) error {
	fmt.Printf("%v: db init ...\n", time.Now())
	afile, err := NewAFile(file)
	if err != nil {
		return err
	}
	db.af = afile
	db.tr, err = genTrie(db.af.fp)
	if err != nil {
		return err
	}
	fmt.Printf("%v: db init finish...\n", time.Now())
	return nil
}

func genTrie(file *os.File) (*patricia.Trie, error) {
	tr := &patricia.Trie{}
	for {
		// t1 := time.Now()
		key, vPos, err := readKey(file)
		//fmt.Printf("read key time: %v\n", time.Now().Sub(t1))
		if err != nil {
			return nil, err
		} else if key == nil {
			break
		} else {
			//t2 := time.Now()
			go tr.Add(key, vPos)
			//fmt.Printf("trie add time: %v\n", time.Now().Sub(t2))
		}
	}
	return tr, nil
}

func readKey(file *os.File) ([]byte, uint64, error) {
	ks := make([]byte, 1)
	if ret, err := file.Read(ks); ret < 0 {
		return nil, 0, err
	} else if ret == 0 || ks[0] == 0 {
		return nil, 0, nil
	}
	key := make([]byte, ks[0])
	if ret, err := file.Read(key); ret < 0 {
		return nil, 0, err
	} else if ret == 0 {
		return nil, 0, nil
	}
	key = bytes.TrimSpace(key)
	vs := make([]byte, 1)
	if ret, err := file.Read(vs); ret < 0 {
		return nil, 0, err
	} else if ret == 0 || vs[0] == 0 {
		return nil, 0, nil
	}
	offset, err := file.Seek(int64(vs[0]), 1)
	if err != nil {
		return nil, 0, err
	}
	vSize := uint64(vs[0])
	vOff := uint64(offset) - vSize
	vPos := (vOff << 32) | vSize

	return key, vPos, nil
}

func (db *DB) Get(key []byte) ([]byte, error) {
	vPos := db.tr.Find(key)
	if vPos == 0 {
		fmt.Printf("vPos is 0, no key in trie\n")
		return nil, nil
	}
	vSize := uint32(vPos & 0xffffffff)
	vOff := uint32((vPos >> 32) & 0xffffffff)
	return db.af.Read(vOff, vSize)
}

func (db *DB) Put(key []byte, value []byte) {
}

func (db *DB) Close() {
	db.af.fp.Close()
}
