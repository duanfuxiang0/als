package alittlestore

import (
	"alittlestore/patricia"
	"os"
)

type DB struct {
	tr *patricia.Trie
	af *AFile
}

func (db *DB) Init(file string) error {
	afile, err := NewAFile(file)
	if err != nil {
		return err
	}
	db.af = afile
	db.tr, err = genTrie(db.af.fp)
	if err != nil {
		return err
	}
	return nil
}

func genTrie(file *os.File) (*patricia.Trie, error) {
	tr := &patricia.Trie{}
	for {
		key, vPos, err := readKey(file)
		if err != nil {
			return nil, err
		} else if key == nil {
			break
		} else {
			tr.Add(key, vPos)
		}
	}
	return tr, nil
}

func readKey(file *os.File) ([]byte, uint64, error) {
	ks := make([]byte, 1)
	//fPos, _ := file.Seek(0, 1)
	//fmt.Println(fPos)
	if ret, err := file.Read(ks); ret < 0 {
		return nil, 0, err
	} else if ret == 0 {
		return nil, 0, nil
	}
	key := make([]byte, ks[0])
	if ret, err := file.Read(key); ret < 0 {
		return nil, 0, err
	} else if ret == 0 {
		return nil, 0, nil
	}
	vs := make([]byte, 1)
	if ret, err := file.Read(vs); ret < 0 {
		return nil, 0, err
	} else if ret == 0 {
		return nil, 0, nil
	}
	offset, _ := file.Seek(int64(vs[0]), 1)
	vSize := uint64(vs[0])
	vOff := uint64(offset - int64(vs[0]))
	vPos := (vOff << 32) | vSize

	return key, vPos, nil
}

func (db *DB) Get(key []byte) ([]byte, error) {
	vPos := db.tr.Find(key)
	if vPos == 0 {
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
