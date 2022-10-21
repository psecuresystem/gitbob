package helpers

import (
	"fmt"
	"strings"
)

func CommitChanges(message string) {
	treeDetails := TransverseTree(".")
	commitData := fmt.Sprintf(`tree %s
	parent %s
	%s`, treeDetails, GetRef("HEAD"), message)
	commitId := HashObject(commitData, "commit")
	SetRef("HEAD",commitId)
	fmt.Printf("Successfully Commited Commit id %s\n", commitId)
}

func GetCommits(from string) string {
	currentCommit := GetRef(from)
	commitHistory := []string{currentCommit}
	return strings.Join(transverseCommitTree(commitHistory), "")
}

func transverseCommitTree(commitHistory []string) []string {
	lastCommit := commitHistory[len(commitHistory) - 1]
	commit := DecodeCommit(lastCommit)
	if commit.Parent == "" || commit.Parent == " " {
		return commitHistory
	}
	fmt.Printf("%s  ->  %s\n", lastCommit, commit.Message)
	commitHistory = append(commitHistory, commit.Parent)

	return transverseCommitTree(commitHistory)
}