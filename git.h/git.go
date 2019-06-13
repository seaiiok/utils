package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"os/exec"
	"utils/fmtx"
)

func main() {
	ExecCommand("chcp", "65001")
	fmtx.Println(fmtx.Ok, "push code...")

	output := ExecCommand("git","add -A")
	fmtx.Println(fmtx.Info, output)

	fmtx.Println(fmtx.Ok, "please input comments...")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	output = ExecCommand("git","commit -m"+input.Text())
	fmtx.Println(fmtx.Info, output)

	output = ExecCommand("git","push origin master")
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
