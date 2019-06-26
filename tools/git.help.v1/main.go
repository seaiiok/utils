package main

import (
	"bufio"
	"os"
	"utils"
)

const (
	meGithub  = "github.com/seaiiok/utils"
	repAlias  = "origin"
	userName  = "seaiiok"
	userEmail = "seaii@qq.com"
	remoteRep = "git@github.com:seaiiok/utils.git"
)

type Git struct {
	meGithub  string
	repAlias  string
	userName  string
	userEmail string
	remoteRep string
}

func main() {

	git := &Git{
		meGithub:  meGithub,
		repAlias:  repAlias,
		userName:  userName,
		userEmail: userEmail,
		remoteRep: remoteRep,
	}
	git.GitLoopHelp()
}

//GitLoopHelp Git帮助
func (g *Git) GitLoopHelp() {

	ii := new(utils.Utils)

	ii.Cmd.ExecCommand("chcp", "65001")

	ii.Cmd.ExecCommand("git", "config", "--global", "user.name", g.userName)
	ii.Cmd.ExecCommand("git", "config", "--global", "user.email", g.userEmail)
	ii.Cmd.ExecCommand("git", "remote", "add", g.repAlias, g.remoteRep)

	for {
		ii.Print.Println(1, "cmd help input push...")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		if input.Text() == "push" {
			ii.Print.Println(1, "push code...")

			output := ii.Cmd.ExecCommand("git", "add", "-A")
			ii.Print.Println(7, output)

			ii.Print.Println(1, "please input comments...")
			input := bufio.NewScanner(os.Stdin)
			input.Scan()

			if input.Text() == "" {
				output = ii.Cmd.ExecCommand("git", "commit", "-m", "update...")
			} else {
				output = ii.Cmd.ExecCommand("git", "commit", "-m", input.Text())
			}

			ii.Print.Println(7, output)

			output = ii.Cmd.ExecCommand("git", "push", g.repAlias, "master")
			ii.Print.Println(7, output)

			output = ii.Cmd.ExecCommand("gopm", "get", "-u", g.meGithub)
			ii.Print.Println(7, output)

			output = ii.Cmd.ExecCommand("gopm", "get", "-g", g.meGithub)
			ii.Print.Println(7, output)

		} else {
			os.Exit(0)
		}
	}
}
