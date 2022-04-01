package service

// todo 进入容器命令行
/*import (
	"kubernetes-admin-backend/client"
	"kubernetes-admin-backend/terminal/websocket"
	"fmt"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"net/http"
)

var (
	cmd = []string{"/bin/sh"}
)

func TerminalPod(w http.ResponseWriter, r *http.Request) {
	ptyHandler, err := websocket.NewTerminalSession(w, r, nil)
	fmt.Println("异常，" + err.Error())
	clientSet, _ := client.GetK8SClientSet()
	req := clientSet.CoreV1().RESTClient().Post().
		Resource("pods").
		Name("details-v1-6c9f8bcbcb-8sm9c").
		Namespace("default").SubResource("exec")

	req.VersionedParams(&v1.PodExecOptions{
		Container: "details",
		Command:   cmd,
		Stdin:     !(ptyHandler.Stdin() == nil),
		Stdout:    !(ptyHandler.Stdout() == nil),
		Stderr:    !(ptyHandler.Stderr() == nil),
		TTY:       ptyHandler.Tty(),
	}, scheme.ParameterCodec)
}*/
