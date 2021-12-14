package main

import (
	"fmt"
	"io"
	"os"
)

type safeWriter struct {
	w   io.Writer
	err error
}

func (sw *safeWriter) writerln(s string) {
	if sw.err != nil {
		return
	}
	_, sw.err = fmt.Fprintln(sw.w, s)

}

func proverbs(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()

	sw := safeWriter{w: f}

	sw.writerln("lcd1")
	sw.writerln("lcd2")
	sw.writerln("lcd3")

	return err
}

func main() {
	err := proverbs("prverbs2.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
