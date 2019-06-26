package utils

import (
	"io/ioutil"
	"os/exec"
)

type git struct {}

func (g *git) ExecCommand(name string, arg ...string) (output string) {
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
