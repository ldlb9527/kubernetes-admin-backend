package client

import (
	"golang.org/x/crypto/ssh"
	"time"
)

type SSHClientConfig struct {
	AuthModel string
	HostAddr  string
	User      string
	Password  string
	PublicKey string
	Timeout   time.Duration
}

func NewSSHClient(conf *SSHClientConfig) (*ssh.Client, error) {
	config := &ssh.ClientConfig{
		Timeout:         conf.Timeout,
		User:            conf.User,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //忽略know_hosts检查
	}
	switch conf.AuthModel {
	case "PASSWORD":
		config.Auth = []ssh.AuthMethod{ssh.Password(conf.Password)}
	case "PUBLICKEY":
		signer, err := ssh.ParsePrivateKey([]byte(conf.PublicKey))
		if err != nil {
			return nil, err
		}
		config.Auth = []ssh.AuthMethod{ssh.PublicKeys(signer)}
	}
	c, err := ssh.Dial("tcp", conf.HostAddr, config)
	if err != nil {
		return nil, err
	}
	return c, nil
}
