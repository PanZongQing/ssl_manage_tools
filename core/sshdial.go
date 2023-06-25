package core

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/ssh"
)

type Config struct {
	SshHost     string
	SshUser     string
	SshPassword string
	SshType     string
	SshPort     string
}

func NewSSH(c *Config) {
	config := &ssh.ClientConfig{
		Timeout:         time.Second,
		User:            c.SshUser,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	if c.SshType == "password" {
		config.Auth = []ssh.AuthMethod{ssh.Password(c.SshPassword)}
	}

	addr := fmt.Sprintf("%s:%s", c.SshHost, c.SshPort)
	sshClient, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		fmt.Println("创建ssh client失败", err)
	}
	defer sshClient.Close()

	session, err := sshClient.NewSession()
	if err != nil {
		fmt.Println("创建session失败", err)
	}
	defer session.Close()

	combo, err := session.CombinedOutput("sudo nginx -t && sudo nginx -s reload")
	if err != nil {
		fmt.Println("远程执行cmd失败", err)

	}
	log.Println("命令输出", string(combo))
}
