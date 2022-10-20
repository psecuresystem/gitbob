package helpers

import (
	"bytes"
	"fmt"
	"os"
)

func WriteFile(path string, oid string) {
	content, _ := os.ReadFile(fmt.Sprintf(".gitb/objects/%s", oid))
	for {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			os.WriteFile(path, bytes.Split(content, []byte{0x00})[1], 0644)
			return
		}
		
	}
}