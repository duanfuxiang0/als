package main

import (
	"alittlestore"
	"fmt"
)

func main() {
	db := &alittlestore.DB{}
	if err := db.Init("db0.data"); err != nil {
		fmt.Printf("err: %v", err)
	}
	val, err := db.Get([]byte("ZamCToZvPynaE"))
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else if val == nil {
		fmt.Println("no this key")
	} else {
		fmt.Printf("val:%s\n", val)
	}
}
