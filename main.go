package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

// Commit :
type Commit map[string]string

func main() {
	out, err := exec.Command("git", "log", "--format=%an %ae %H %aD Subject:%s").Output()
	if err != nil {
		fmt.Println("error occured")
		fmt.Printf("%s", err)
	}

	commits := strings.Split(string(out), "\n")
	commitsArray := make([]string, len(commits)-1)

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

			data, err := json.Marshal(commitStructure)
			if err != nil {
				panic(err)
			}
			commitsArray[i] = string(data)
		}

		commitJSON, _ := json.Marshal(commitsArray)

		fmt.Printf(string(commitJSON))
	}
}
