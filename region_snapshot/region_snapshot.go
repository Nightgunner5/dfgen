package main

import (
	"bufio"
	"compress/gzip"
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

	g, err := gzip.NewReader(f)
	handle(err)
	defer g.Close()

	r := dfgen.Reader{bufio.NewReader(g)}

	r.ReadLong() // version number (1404)
	r.ReadLong() // compression state (0)

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

	var buf [17 * 17]byte
	_, err = io.ReadFull(r.R, buf[:])
	handle(err)
	for i := 0; i < 17; i++ {
		fmt.Println("Unk9:", i, buf[i*17:][:17])
	}

	x := hex.Dumper(os.Stdout)
	defer x.Close()

	_, err = io.Copy(x, r.R)
	handle(err)
}
