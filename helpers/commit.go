package helpers

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

type Commit struct {
	Tree string
	Parent string
	Message string
}

func DecodeCommit(oid string) *Commit {
	commitContent, _ := os.ReadFile(fmt.Sprintf(".gitb/objects/%s", strings.TrimSpace(oid)))
	splitCommitContent := bytes.Split(commitContent, []byte{0x00})

	if len(splitCommitContent) == 1 {
		return &Commit{Tree: "", Parent: "", Message: ""}
	}
	commitLines := strings.Split(string(splitCommitContent[1]), "\n")

	return &Commit{
		Tree: strings.Split(commitLines[0], "tree ")[1],
		Parent: strings.Split(commitLines[1], "parent ")[1], 
		Message: strings.TrimSpace(strings.Split(commitLines[2], "message ")[0]),
	}
}

