package main

import (
	"als"
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {

	file, err := os.OpenFile("db0.data", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	bw := bufio.NewWriter(file)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	n := 1024
	for n > 0 {
		key := als.RandStringRunes(r)
		if err := bw.WriteByte(byte(len(key))); err != nil {
			fmt.Printf("err: %v", err)
		}
		if _, err := bw.Write([]byte(key)); err != nil {
			fmt.Printf("err: %v", err)
		}

		val := als.RandStringRunes(r)
		if err := bw.WriteByte(byte(len(val))); err != nil {
			fmt.Printf("err: %v", err)
		}
		if _, err := bw.Write([]byte(val)); err != nil {
			fmt.Printf("err: %v", err)
		}
		fmt.Printf("k: %s v: %s\n", key, val)
		if bw.Available() < 514 {
			if err := bw.Flush(); err != nil {
				fmt.Printf("err: %v", err)
				break
			} else {
				n = n - 1
			}
		}
	}
	bw.Flush()
}