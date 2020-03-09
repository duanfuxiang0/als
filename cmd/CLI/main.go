package main

import (
	"als"
	"fmt"
	"net/rpc"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: miss the key...\n")
		os.Exit(1)
	}
	key := os.Args[1]

	c, err := rpc.DialHTTP("tcp", "127.0.0.1:9527")
	arg := als.GetArgs{Key:[]byte(key)}
	reply := als.GetReply{}
	err = c.Call("ALS.GetKey", &arg, &reply)
	if err != nil {
		fmt.Printf("get key: %s err: %v\n", key, err)
	}
	fmt.Printf("%s\n", string(reply.Val))
}