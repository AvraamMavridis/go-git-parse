package main

import (
	"fmt"
	"os/exec"
)

func main() {
	out, err := exec.Command("git", "log").Output()
	if err != nil {
		fmt.Println("error occured")
		fmt.Printf("%s", err)
	}
	fmt.Printf(string(out))
}
