package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"os/exec"
	"github.com/seaiiok/utils/fmtx"
)

func main() {
	ExecCommand("chcp", "65001")

	for {
		fmtx.Println(fmtx.Ok, "cmd help input push...")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		if input.Text() == "push" {
			fmtx.Println(fmtx.Ok, "push code...")
			GitPush()
		}

	}
}

func GitPush() {
	output := ExecCommand("git", "add", "-A")
	fmtx.Println(fmtx.Info, output)

	fmtx.Println(fmtx.Ok, "please input comments...")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	if input.Text() == "" {
		output = ExecCommand("git", "commit", "-m", "update...")
	} else {
		output = ExecCommand("git", "commit", "-m", input.Text())
	}

	fmtx.Println(fmtx.Info, output)

	output = ExecCommand("git", "push", "origin", "master")
	fmtx.Println(fmtx.Info, output)

	output = ExecCommand("gopm", "get", "-u", "github.com/seaiiok/utils")
	fmtx.Println(fmtx.Info, output)

	output = ExecCommand("gopm", "get", "-g", "github.com/seaiiok/utils")
	fmtx.Println(fmtx.Info, output)
}

func ExecCommand(name string, arg ...string) (output string) {
	cmd := exec.Command(name, arg...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err.Error()
	}

	defer stdout.Close()

	err = cmd.Start()
	if err != nil {
		return err.Error()
	}

	opBytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		return err.Error()
	} else {
		return string(opBytes)
	}
}
