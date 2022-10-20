package helpers

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
)

func HashObject(fileContent string, typeOf string) string {
	hasher := sha1.New()
	hasher.Write([]byte(fileContent))
	hashedFileName := hex.EncodeToString(hasher.Sum(nil))

	pathName := fmt.Sprintf(".gitb/objects/%s", hashedFileName)
	file, _ := os.Create(pathName)

	file.Write(bytes.Join([][]byte{[]byte(typeOf)[:], []byte(fileContent)}, []byte{0x00}))
	fmt.Printf("Successfully written object %s \n", hashedFileName)
	return hashedFileName
}