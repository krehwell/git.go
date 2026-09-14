package main

import (
	"bufio"
	"crypto/sha1"
	"fmt"
	"os"
)

func hashBlob(content []byte) string {
	header := []byte(fmt.Sprintf("blob %d\x00", len(content)))
	// TODO: combined := append(header, content...)
	// TODO: h := sha1.Sum(combined)
	// TODO: return fmt.Sprintf("%x", h)
	_ = header
	return ""
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1<<20), 1<<20)
	for sc.Scan() {
		fmt.Println(hashBlob(sc.Bytes()))
	}
}
