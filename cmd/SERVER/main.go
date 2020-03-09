package main

import (
	"als"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: miss dbfile...\n")
		os.Exit(1)
	}
	store := als.MakeServer(os.Args[1])
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
	for {
		select {
		case <-c:
			als.Clear(store)
			fmt.Printf("als exit...")
			return
		case <-store.Quit:
			als.Clear(store)
			fmt.Printf("als exit...")
			return
		}
	}
}
