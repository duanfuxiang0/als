package main

import (
	"als"
	"fmt"
	"net/rpc"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: miss db addr...\n")
		os.Exit(1)
	}
	db := os.Args[1]
	key := os.Args[2]

	c, err := rpc.DialHTTP("tcp", db)
	arg := als.GetArgs{Key: []byte(key)}
	reply := als.GetReply{}
	err = c.Call("ALS.GetKey", &arg, &reply)
	if err != nil {
		fmt.Printf("get key: %s err: %v\n", key, err)
	}
	fmt.Printf("%s\n", string(reply.Val))
}
