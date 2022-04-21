module kubernetes-admin-backend

go 1.16

require (
	github.com/gin-gonic/gin v1.7.7
	github.com/gorilla/sessions v1.2.1
	github.com/gorilla/websocket v1.5.0
	github.com/imdario/mergo v0.3.7 // indirect
	github.com/prometheus/client_golang v1.12.1
	golang.org/x/crypto v0.0.0-20201002170205-7f63de1d35b0
	k8s.io/api v0.20.1
	k8s.io/apimachinery v0.20.1
	k8s.io/client-go v0.20.1
	k8s.io/klog v1.0.0
	sigs.k8s.io/yaml v1.2.0
)
