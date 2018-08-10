package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Commit : Map of strings
type Commit map[string]string

func main() {
	out, err := exec.Command("git", "log", "--format=%an %ae %H %aD Subject:%s").Output()
	if err != nil {
		fmt.Println("error occured")
		fmt.Printf("%s", err)
	}

	commits := strings.Split(string(out), "\n")
	commitsArray := make([]Commit, len(commits)-1)

	for i, commit := range commits {
		var commitInfo = strings.Split(string(commit), " ")

		if len(commitInfo) > 3 {
			var authorName = commitInfo[0]
			var authorEmail = commitInfo[1]
			var commitHash = commitInfo[2]

			commitStructure := Commit{
				"authorName":  authorName,
				"authorEmail": authorEmail,
				"commitHash":  commitHash,
			}

			commitsArray[i] = commitStructure
		}

		file, _ := os.Create("git.json")
		commitJSON, _ := json.Marshal(commitsArray)
		defer file.Close()

		fmt.Fprintf(file, string(commitJSON))
	}
}
