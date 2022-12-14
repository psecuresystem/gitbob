package helpers

import (
	"fmt"
	"os"
)

func SetRef(ref string, oid string) {
	file, _ := os.Create(fmt.Sprintf(".gitb/%s", ref))
	file.WriteString(oid)
	fmt.Printf("Set Ref %s to %s\n", ref, oid)
}

func GetRef(from string) string {
	if from  == "@" {
		from = "HEAD"
	}
	refs_to_try := []string{
		fmt.Sprintf(".gitb/%s", from),
		fmt.Sprintf(".gitb/refs/%s", from),
		fmt.Sprintf(".gitb/refs/tags/%s", from),
		fmt.Sprintf(".gitb/refs/heads/%s", from),
	}

	for _, ref_to_try := range refs_to_try {
		file, err := os.ReadFile(ref_to_try)
		if os.IsNotExist(err) {
			continue
		}
		return string(file)
	}

	if len(from) == 40 {
		return from
	}
	return ""
}

func IterRefs() {}