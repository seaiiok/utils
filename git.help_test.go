package utils

import (
	"bufio"
	"os"
	"testing"
)

func TestGitHelp(t *testing.T) {
	const (
		git       = "git"
		gopm      = "gopm"
		meGithub  = "github.com/seaiiok/utils"
		repAlias  = "origin"
		userName  = "seaiiok"
		userEmail = "seaii@qq.com"
		remoteRep = "git@github.com:seaiiok/utils.git"
	)
	ii := new(Utils)

	ii.Git.ExecCommand("chcp", "65001")

	ii.Git.ExecCommand(git, "config", "--global", "user.name", userName)
	ii.Git.ExecCommand(git, "config", "--global", "user.email", userEmail)
	ii.Git.ExecCommand(git, "remote", "add", repAlias, remoteRep)

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

	output = ii.Git.ExecCommand(git, "push", repAlias, "master")
	ii.Print.Println(7, output)

	output = ii.Git.ExecCommand(gopm, "get", "-u", meGithub)
	ii.Print.Println(7, output)

	output = ii.Git.ExecCommand(gopm, "get", "-g", meGithub)
	ii.Print.Println(7, output)

}

func TestGitLoopHelp(t *testing.T) {
	const (
		git       = "git"
		gopm      = "gopm"
		meGithub  = "github.com/seaiiok/utils"
		repAlias  = "origin"
		userName  = "seaiiok"
		userEmail = "seaii@qq.com"
		remoteRep = "git@github.com:seaiiok/utils.git"
	)
	ii := new(Utils)

	ii.Git.ExecCommand("chcp", "65001")

	ii.Git.ExecCommand(git, "config", "--global", "user.name", userName)
	ii.Git.ExecCommand(git, "config", "--global", "user.email", userEmail)
	ii.Git.ExecCommand(git, "remote", "add", repAlias, remoteRep)

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

			output = ii.Git.ExecCommand(git, "push", repAlias, "master")
			ii.Print.Println(7, output)

			output = ii.Git.ExecCommand(gopm, "get", "-u", meGithub)
			ii.Print.Println(7, output)

			output = ii.Git.ExecCommand(gopm, "get", "-g", meGithub)
			ii.Print.Println(7, output)

		} else {
			os.Exit(0)
		}
	}
}
