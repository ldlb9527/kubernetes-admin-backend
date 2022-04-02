package terminal

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"golang.org/x/crypto/ssh"
	"io"
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

type Turn struct {
	StdinPipe io.WriteCloser
	Session   *ssh.Session
	WsConn    *websocket.Conn
}

func NewTurn(wsConn *websocket.Conn, sshClient *ssh.Client) (*Turn, error) {
	sess, err := sshClient.NewSession()
	if err != nil {
		return nil, err
	}

	stdinPipe, err := sess.StdinPipe()
	if err != nil {
		return nil, err
	}

	turn := &Turn{StdinPipe: stdinPipe, Session: sess, WsConn: wsConn}
	sess.Stdout = turn
	sess.Stderr = turn

	modes := ssh.TerminalModes{
		ssh.ECHO:          1,     // disable echo
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}
	if err := sess.RequestPty("xterm", 150, 30, modes); err != nil {
		return nil, err
	}
	if err := sess.Shell(); err != nil {
		return nil, err
	}

	return turn, nil
}

func (t *Turn) Write(p []byte) (n int, err error) {
	writer, err := t.WsConn.NextWriter(websocket.TextMessage)
	if err != nil {
		return 0, err
	}
	defer writer.Close()
	fmt.Println("Write:" + string(p))
	return writer.Write(p)
}
func (t *Turn) Close() error {
	fmt.Println("Close()")
	if t.Session != nil {
		t.Session.Close()
	}

	return t.WsConn.Close()
}

func (t *Turn) Read(p []byte) (n int, err error) {
	for {
		msgType, reader, err := t.WsConn.NextReader()
		if err != nil {
			return 0, err
		}
		if msgType != websocket.TextMessage {
			continue
		}
		fmt.Println("Write:" + string(p))
		return reader.Read(p)
	}
}

func (t *Turn) LoopRead(context context.Context) error {
	for {
		select {
		case <-context.Done():
			return errors.New("LoopRead exit")
		default:
			_, wsData, err := t.WsConn.ReadMessage()
			fmt.Println("本地输入：" + string(wsData))
			if err != nil {
				return fmt.Errorf("reading webSocket message err:%s", err)
			}
			body := decode(wsData[1:])
			fmt.Println("body:" + string(body))
			body = wsData
			fmt.Println("body:" + string(body))
			if _, err := t.StdinPipe.Write(body); err != nil {
				return fmt.Errorf("StdinPipe write err:%s", err)
			}

			/*switch wsData[0] {
			case MsgResize:
				var args Resize
				err := json.Unmarshal(body, &args)
				if err != nil {
					return fmt.Errorf("ssh pty resize windows err:%s", err)
				}
				if args.Columns > 0 && args.Rows > 0 {
					if err := t.Session.WindowChange(args.Rows, args.Columns); err != nil {
						return fmt.Errorf("ssh pty resize windows err:%s", err)
					}
				}
			case MsgData:
				if _, err := t.StdinPipe.Write(body); err != nil {
					return fmt.Errorf("StdinPipe write err:%s", err)
				}
				if _, err := logBuff.Write(body); err != nil {
					return fmt.Errorf("logBuff write err:%s", err)
				}
			}*/
		}
	}
}

func (t *Turn) SessionWait() error {
	if err := t.Session.Wait(); err != nil {
		return err
	}
	return nil
}

func decode(p []byte) []byte {
	decodeString, _ := base64.StdEncoding.DecodeString(string(p))
	return decodeString
}
