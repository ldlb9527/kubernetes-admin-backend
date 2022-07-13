package test

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"testing"
)

func TestAddCluster(t *testing.T) {
	bytes, err := ioutil.ReadFile("../config/.kube/config")
	if err != nil {
		fmt.Println("ReadFile," + err.Error())
	}
	config, err := clientcmd.NewClientConfigFromBytes(bytes)
	if err != nil {
		fmt.Println("NewClientConfigFromBytes," + err.Error())
	}
	clientConfig, err := config.ClientConfig()
	if err != nil {
		fmt.Println("ClientConfig," + err.Error())
	}
	clientSet, err := kubernetes.NewForConfig(clientConfig)
	if err != nil {
		fmt.Println("NewForConfigorConfigFile," + err.Error())
	}
	version, err := clientSet.ServerVersion()
	if err != nil {
		fmt.Println("ServerVersion," + err.Error())
	}
	fmt.Println(version)

}

func TestFileToBytes(t *testing.T) {
	file, err := os.Open("../config/.kube/config")
	if err != nil {
		fmt.Println("Open," + err.Error())
	}
	defer file.Close()
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		fmt.Println("io.Copy," + err.Error())
	}
	config, err := clientcmd.NewClientConfigFromBytes(buf.Bytes())
	if err != nil {
		fmt.Println("NewClientConfigFromBytes," + err.Error())
	}
	clientConfig, err := config.ClientConfig()
	if err != nil {
		fmt.Println("ClientConfig," + err.Error())
	}
	clientSet, err := kubernetes.NewForConfig(clientConfig)
	if err != nil {
		fmt.Println("NewForConfigorConfigFile," + err.Error())
	}
	version, err := clientSet.ServerVersion()
	if err != nil {
		fmt.Println("ServerVersion," + err.Error())
	}
	fmt.Println(version)

}
