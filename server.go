package als

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type ALS struct {
	db *DB
	Quit chan struct{}
}

type GetArgs struct {
	Key []byte
}

type GetReply struct {
	Val []byte
}

func (als *ALS) GetKey(args *GetArgs, reply *GetReply) error {
	val, err := als.db.Get(args.Key)
	if err != nil {
		return err
	}
	*reply = GetReply{Val: val}
	return nil
}

//type SetArgs struct {
//	key []byte
//}
//
//type SetReply struct {
//	val []byte
//}
//
//func (als *ALS) SetKey() {
//
//}

func (als *ALS) server() {
	rpc.Register(als)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":9527")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}

func Clear(als *ALS) {
	als.db.Close()
}

func MakeServer(dbfile string) *ALS {
	quit := make(chan struct{})
	als := ALS{db: &DB{}, Quit: quit}
	if err := als.db.Init(dbfile); err != nil {
		fmt.Printf("db init error: %v", err)
	}
	als.server()
	return &als
}
