package main

import (
	"als"
	"fmt"
	"os"
	"os/signal"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: miss dbfile...\n")
		os.Exit(1)
	}
	store := als.MakeServer(os.Args[1])

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	for {
		select {
		case <-sigs:
		case <-store.Quit:
			als.Clear(store)
			fmt.Printf("als exit...")
			return
		}
	}
}
