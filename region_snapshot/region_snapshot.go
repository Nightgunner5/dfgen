package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"github.com/Nightgunner5/dfgen"
	"io"
	"os"
)

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	f, err := os.Open(os.Args[1])
	handle(err)
	defer f.Close()

	r, err := dfgen.NewReader(bufio.NewReaderSize(f, 0x10000))
	handle(err)

	n, err := r.ReadLong()
	handle(err)
	fmt.Println("Year:", n)

	length, err := r.ReadLong()
	handle(err)
	fmt.Println("\nUnk2:\nLength", length)
	for i := 0; i < int(length); i++ {
		n, err = r.ReadLong() // are these ALWAYS the first N natural numbers?
		handle(err)
		fmt.Println(i, n)
	}

	length, err = r.ReadLong() // is this ALWAYS the same as the previous length?
	handle(err)
	fmt.Println("\nUnk3:\nLength", length)
	for i := 0; i < int(length); i++ {
		n, err = r.ReadLong()
		handle(err)
		fmt.Println(i, n)
	}

	b, err := r.ReadByte() // is this ALWAYS 1?
	handle(err)
	fmt.Println("Unk4:", b)

	s, err := r.ReadShort() // is this ALWAYS 2?
	handle(err)
	fmt.Println("Unk5:", s)

	s, err = r.ReadShort() // is this ALWAYS 2?
	handle(err)
	fmt.Println("Unk6:", s)

	length, err = r.ReadLong() // is this ALWAYS 10?
	handle(err)
	fmt.Println("\nUnk7:\nLength", length)
	for i := 0; i < int(length); i++ {
		n, err = r.ReadLong()
		handle(err)
		fmt.Println(i, n)
	}

	s, err = r.ReadShort() // is this ALWAYS 2?
	handle(err)
	fmt.Println("Unk8:", s)

	x := hex.Dumper(os.Stdout)
	defer x.Close()

	_, err = io.Copy(x, r.R)
	handle(err)
}
