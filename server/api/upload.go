package api

import (
	"log"
	"os"
	"ssl_manage/core"

	"github.com/povsister/scp"
)

func CopyFileToRemote(ipaddr, username, password string, filename string) error {
	sshConf := scp.NewSSHConfigFromPassword(username, password)
	scpClient, err := scp.NewClient(ipaddr, sshConf, &scp.ClientOption{
		Sudo: true,
	})
	if err != nil {
		log.Fatalf("err massage : %s", err)
	}
	defer scpClient.Close()

	localpath := "./uploaddir/" + filename
	sourcepath := "/opt/" + filename

	err = scpClient.CopyFileToRemote(localpath, sourcepath, &scp.FileTransferOption{
		Perm: 0644,
	})

	return err
}

func CopyRemoteToLocal(ipaddr, username, password string) error {
	sshConf := scp.NewSSHConfigFromPassword(username, password)
	scpClient, err := scp.NewClient(ipaddr, sshConf, &scp.ClientOption{})
	if err != nil {
		log.Fatalf("err massage : %s", err)
	}
	defer scpClient.Close()

	confpath := "./confpath"

	if core.PathExists(confpath) == false {
		err := os.Mkdir("./confpath", os.ModePerm)
		if err != nil {
			return nil
		}
	}

	localpath := "./confpath"
	remotepath := "/etc/nginx/conf.d"
	err = scpClient.CopyDirFromRemote(remotepath, localpath, &scp.DirTransferOption{})

	return err
}
