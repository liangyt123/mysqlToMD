package main

import (
	"errors"
	"fmt"
	"os"
)

const mdHelpString = "-h 数据库地址 \n" +
	"-u 用户名 \n" +
	"-p 密码 \n" +
	"-P 端口 \n" +
	"-d 数据库名 \n" +
	"-t 表名 \n" +
	"--help 帮助 \n" +
	"说明:\n" +
	"  1.-u -p -d 必填参数.\n" +
	"  2.命令后面要打一个空格，然后跟参数 \n" +
	"  如：-h 127.0.0.1 -u root -p root -d database -P 3306  \n"

func PrintCmdHelp() {
	fmt.Println(mdHelpString)
}

//解释参数
func ParseCommnd() (*ParamStruct, error) {
	param := &ParamStruct{}
	//将cmd参数分割
	//cmds := strings.Split(strcmd, "")
	cmds := os.Args
	if len(cmds) == 2 && cmds[1] == "--help" {
		PrintCmdHelp()
	}
	//fmt.Println("控制台输入：", cmds)

	if len(cmds) < 6 {
		fmt.Println("参数不全")
		PrintCmdHelp()
		return nil, fmt.Errorf("no error")
	}

	for i := 1; i < len(cmds); i = i + 2 {
		if i == 0 {
			continue
		}
		if cmds[i] == "-u" {
			param.User = cmds[i+1]
		} else if cmds[i] == "-p" {
			param.Password = cmds[i+1]
		} else if cmds[i] == "-d" {
			param.Dasebase = cmds[i+1]
		} else if cmds[i] == "-t" {
			param.Table = cmds[i+1]
		} else if cmds[i] == "-h" {
			param.Host = cmds[i+1]
		} else if cmds[i] == "-P" {
			param.Port = cmds[i+1]
		} else if cmds[i] == "--path" {
			param.Path = cmds[i+1]
		} else {
			fmt.Println(cmds[i], "is Illegal parameters!")
			PrintCmdHelp()
			return nil, errors.New("error param")
		}
	}
	return param, nil
}
