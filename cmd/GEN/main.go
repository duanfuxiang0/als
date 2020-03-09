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
	file, err := os.OpenFile("db1.data", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	bw := bufio.NewWriterSize(file, 80960)
	defer func() {
		bw.Flush()
	}()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	n := 1024 * 1024 * 275941
	for n > 0 {
		ks, key := als.RandBytes(r, 50)
		if err := bw.WriteByte(byte(ks)); err != nil {
			fmt.Printf("err: %v\n", err)
		}
		if _, err := bw.Write(key); err != nil {
			fmt.Printf("err: %v\n", err)
		}
		vs, val := als.RandBytes(r, 200)
		if err := bw.WriteByte(byte(vs)); err != nil {
			fmt.Printf("err: %v\n", err)
		}
		if _, err := bw.Write(val); err != nil {
			fmt.Printf("err: %v\n", err)
		}
		fmt.Printf("%s\n", string(key))
		if bw.Available() < 514 {
			if err := bw.Flush(); err != nil {
				fmt.Printf("err: %v\n", err)
				break
			} else {
				n = n - 1
			}
		}
	}
}