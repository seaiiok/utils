package utils

import (
	"bufio"
	"os"
	"testing"
	"time"
)

func TestGit(t *testing.T) {
	const (
		git      = "git"
		gopm     = "gopm"
		meGithub = "github.com/seaiiok/utils"
		RepAlias = "origin"
	)
	ii := new(Utils)
	ii.Git.UserName = "seaiiok"
	ii.Git.UserEmail = "seaii@qq.com"
	ii.Git.RemoteRep = "git@github.com:seaiiok/utils.git"

	ii.Git.ExecCommand("chcp", "65001")

	ii.Git.ExecCommand(git, "config", "--global", "user.name", ii.Git.UserName)
	ii.Git.ExecCommand(git, "config", "--global", "user.email", ii.Git.UserEmail)
	ii.Git.ExecCommand(git, "remote", "add", RepAlias, ii.Git.RemoteRep)

	for {
		ii.Print.Println(1, "cmd help input push...")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		if input.Text() == "push" {
			ii.Print.Println(1, "push code...")

			output := ii.Git.ExecCommand(git, "add", "-A")
			ii.Print.Println(7, output)

			ii.Print.Println(1, "please input comments...")
			input := bufio.NewScanner(os.Stdin)
			input.Scan()

			if input.Text() == "" {
				output = ii.Git.ExecCommand(git, "commit", "-m", "update...")
			} else {
				output = ii.Git.ExecCommand(git, "commit", "-m", input.Text())
			}

			ii.Print.Println(7, output)

			output = ii.Git.ExecCommand(git, "push", RepAlias, "master")
			ii.Print.Println(7, output)

			output = ii.Git.ExecCommand(gopm, "get", "-u", meGithub)
			ii.Print.Println(7, output)

			output = ii.Git.ExecCommand(gopm, "get", "-g", meGithub)
			ii.Print.Println(7, output)

		} else {
			// os.Exit(0)
		}
		time.Sleep(1 * time.Second)
	}

}
