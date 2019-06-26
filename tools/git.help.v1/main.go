package main

import (
	"bufio"
	"os"
	"strings"
	"utils"
)

//自己的github信息
//先配置好自己的git ssh 公钥
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
	// 执行git帮助
	git.GitLoopHelp()
}

// GitLoopHelp 执行git相关命令
// git add -A
// git commit -m ""
// git push remoteRep master
// 如需pull,请手动
func (g *Git) GitLoopHelp() {

	ii := new(utils.Utils)

	//避免控制台显示乱码，临时采用UTF-8
	ii.Cmd.ExecCommand("chcp", "65001")

	//主机
	host, err := os.Hostname()
	if err != nil {
		host = "administrator"
	}
	host = strings.ToLower(host)

	//配置git
	//首次使用手动 git init
	ii.Cmd.ExecCommand("git", "config", "--global", "user.name", g.userName)
	ii.Cmd.ExecCommand("git", "config", "--global", "user.email", g.userEmail)
	ii.Cmd.ExecCommand("git", "remote", "add", g.repAlias, g.remoteRep)

	//循环任务
	//当前仅支持命令:
	//push  推送任务
	//exit  退出git帮助
	ii.Print.Println(13, "start git help...")
	for {
		ii.Print.Println(11, "cmd help input -push...")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()

		if strings.ToLower(strings.TrimSpace(input.Text())) == "push" {
			ii.Print.Println(11, "push code...")

			output := ii.Cmd.ExecCommand("git", "add", "-A")
			ii.Print.Println(8, output)

			ii.Print.Println(11, "please input comments...")
			input := bufio.NewScanner(os.Stdin)
			input.Scan()

			if input.Text() == "" {
				output = ii.Cmd.ExecCommand("git", "commit", "-m", "update by "+host)
			} else {
				output = ii.Cmd.ExecCommand("git", "commit", "-m", strings.ToLower(strings.TrimSpace(input.Text())))
			}

			ii.Print.Println(8, output)

			output = ii.Cmd.ExecCommand("git", "push", g.repAlias, "master")
			ii.Print.Println(8, output)

			//更新本地代码库
			output = ii.Cmd.ExecCommand("gopm", "get", "-u", g.meGithub)
			ii.Print.Println(8, output)

			output = ii.Cmd.ExecCommand("gopm", "get", "-g", g.meGithub)
			ii.Print.Println(8, output)

		} else if strings.ToLower(strings.TrimSpace(input.Text())) == "exit" {
			os.Exit(0)
		}
	}
}
