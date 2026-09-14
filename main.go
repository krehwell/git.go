package main

import (
	"bufio"
	"crypto/sha1"
	"fmt"
	"os"
)

func hashBlob(content []byte) string {
	header := []byte(fmt.Sprintf("blob %d\x00", len(content)))
	combined := append(header, content...)
	h := sha1.Sum(combined)
	return fmt.Sprintf("%x", h)
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1<<20), 1<<20)
	for sc.Scan() {
		fmt.Println(hashBlob(sc.Bytes()))
	}
}
