package helpers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func ReadTree(oid string, rootDir string) string {
	if !FileExists(fmt.Sprintf(".gitb/objects/%s", oid)) {
		fmt.Println("OID Doesnt exist")
		return ""
	}
	fileContent, _ := os.ReadFile(fmt.Sprintf(".gitb/objects/%s", oid))
	fileDetails := bytes.Split(fileContent, []byte{0x00})
	if string(fileDetails[0]) != "tree" {
		fmt.Println("Make sure the oid is that of a tree")
		return ""
	}
	deleteTree(".", []string{})
	// Loop through file Contents
	for _, element := range strings.Split(string(fileDetails[1]), "\n") {
		line := strings.Split(element, " ")
		if line[0] == "tree" {
			ReadTree(line[1], fmt.Sprintf("%s/%s", rootDir, line[2]))
		} else {
			fileDir := fmt.Sprintf("%s/%s", rootDir, line[2])
			WriteFile(fileDir, line[1])
			fmt.Printf("Successfully Written File %s\n", fileDir)
		}
	}
	return ""
}


func deleteTree(rootDir string, deleted []string) {
	fileInfo, _ := ioutil.ReadDir(rootDir)
	for _, file := range fileInfo {
		if file.Name() == ".gitb" || file.Name() == ".git" {
			continue
		}
		if file.IsDir() {
			deleteTree(fmt.Sprintf("%s/%s", rootDir, file.Name()), deleted)
		} else {
			os.Remove(fmt.Sprintf("%s/%s", rootDir, file.Name()))
			fmt.Printf("Deleting Path %s\n", fmt.Sprintf("%s/%s", rootDir, file.Name()))
			deleted = append(deleted, fmt.Sprintf("%s/%s", rootDir, file.Name()))
		}
	}
}