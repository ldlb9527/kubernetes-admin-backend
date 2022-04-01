package apis

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"kubernetes-admin-backend/client"
	"log"
	"net/http"
	"sync"
	"time"
)

//包级变量,升级器
var upgrader = websocket.Upgrader{}

func init() {
	//初始化
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		// 解决跨域问题
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
}

// VisitorWebsocketServer https://github.com/widaT/webssh  websocket连接实现webssh
func VisitorWebsocketServer(c *gin.Context) {
	wsConn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("upgrade error:", err)
		return
	}
	defer wsConn.Close()

	config := &client.SSHClientConfig{
		Timeout:   time.Second * 5,
		HostAddr:  "xxx.xxx.xxx.xxx:22",
		User:      "*****",
		Password:  "*****",
		AuthModel: "PASSWORD",
	}
	sshClient, err := client.NewSSHClient(config)
	if err != nil {
		wsConn.WriteControl(websocket.CloseMessage,
			[]byte(err.Error()), time.Now().Add(time.Second))
		return
	}
	defer sshClient.Close()

	turn, err := client.NewTurn(wsConn, sshClient)
	if err != nil {
		fmt.Println("NewTurn," + err.Error())
		wsConn.WriteControl(websocket.CloseMessage,
			[]byte(err.Error()), time.Now().Add(time.Second))
		return
	}
	defer turn.Close()

	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		err := turn.LoopRead(ctx)
		if err != nil {
			log.Printf("%#v", err)
		}
	}()
	go func() {
		defer wg.Done()
		err := turn.SessionWait()
		if err != nil {
			log.Printf("%#v", err)
		}
		cancel()
	}()
	wg.Wait()
}
