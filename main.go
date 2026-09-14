package main

import (
	"bufio"
	"bytes"
	"crypto/sha1"
	"fmt"
	"os"
	"strings"
)

func hashBlob(content []byte) string {
	header := []byte(fmt.Sprintf("blob %d\x00", len(content)))
	combined := append(header, content...)
	h := sha1.Sum(combined)
	return fmt.Sprintf("%x", h)
}

var objectTypes = map[string]bool{"blob": true, "tree": true, "commit": true, "tag": true}

func parseBlob(content []byte) string {
	// The separator is the literal text \0, not a NUL byte.
	i := bytes.Index(content, []byte(`\0`))
	if i == -1 {
		return "ERR no NUL separator"
	}
	header, body := content[:i], content[i+2:]

	typ, size, _ := strings.Cut(string(header), " ")
	if !objectTypes[typ] {
		return "ERR unknown type " + typ
	}

	return fmt.Sprintf("type %s\nsize %s\nbody %d", typ, size, len(body))
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1<<20), 1<<20)
	for sc.Scan() {
		fmt.Println(parseBlob(sc.Bytes()))
	}
}
