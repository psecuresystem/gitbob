package helpers

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func TransverseTree(rootDir string) string {
	var files []string
	fileInfo, err := ioutil.ReadDir(rootDir)
	if err != nil {
		fmt.Println("An error")
	}

	for _, file := range fileInfo { 
		fileContent, _ := os.ReadFile(fmt.Sprintf("%s/%s", rootDir, file.Name()))
		if file.Name() == ".gitb" || file.Name() == ".git" {
			continue
		}
		
		if !file.IsDir() {
			fileNameHash := HashObject(string(fileContent), "blob")
			files = append(files, fmt.Sprintf("blob %s %s", fileNameHash, file.Name()))
		} else {
			dirHash := TransverseTree(fmt.Sprintf("%s/%s", rootDir, file.Name()))
			files = append(files, fmt.Sprintf("tree %s %s", dirHash, file.Name()))
		}
	}

	return HashObject(strings.Join(files, "\n"), "tree")
}