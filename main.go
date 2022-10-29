package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("incorrect use of tool")
		fmt.Println("provide commit msg, e.g. commit0r.exe \"COMMIT MSG\"")
		return
	}

	gitAdd()

	msg := os.Args[1]
	commitmsg := getCurrentBranch() + ": " + msg

	gitCommit(commitmsg)
	gitPush()
}

func gitPush() {
	err := exec.Command("git", "push").Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
}

func gitCommit(msg string) {
	err := exec.Command("git", "commit", "-m", msg).Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
}

func getCurrentBranch() string {
	branch, err := exec.Command("git", "branch", "--show-current").Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
	return strings.Split(string(branch), "\n")[0]
}

func gitAdd() {
	err := exec.Command("git", "add", ".").Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
}
