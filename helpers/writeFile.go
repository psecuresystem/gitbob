package helpers

import (
	"bytes"
	"fmt"
	"os"
)

func WriteFile(path string, oid string) {
	content, _ := os.ReadFile(fmt.Sprintf(".gitb/objects/%s", oid))
	os.WriteFile(path, bytes.Split(content, []byte{0x00})[1], 0644)
}