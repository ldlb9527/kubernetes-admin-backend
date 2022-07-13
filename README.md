![](https://img.shields.io/badge/-kubernetes--admin--backend-green)
# k8s可视化管理平台
## 1.项目结构介绍
* `apis`: 控制器,接口访问的入口,文件大多以k8s资源命名
* `config`: 配置文件,程序端口和k8s的config配置
* `middleware`:中间件，配置跨域访问和prometheus监控
* `proto`: 返回给前端的实体，类似于VO对象,Cluster.go用于配置多集群管理
* `router`: 路由层,路由分组和匹配,与apis包中函数绑定
* `service`: 业务处理层
* `terminal`: 提供web界面进入pod和webssh
***
## 2.如何运行(本地运行)
* ①首先你需要一个k8s集群，用k8s master节点下的/root/.kube/config替换当前项目./config/.kube/default/config
* ②如果你的本地与k8s集群在同一内网环境，然后就可以直接运行了
* ③如果你的本地与k8s集群不在同一内网环境(比如集群在云端)，本地运行会报错，需要扩展k8s访问地址,[k8s扩展地址参考](https://blog.csdn.net/marlinlm/article/details/122166105)
* ④扩展地址后,将config放在对应目录`./config/.kube/config`,加载配置文件的代码在client包
* ⑤将config文件的server修改为你扩展的外网ip，然后直接运行
* ⑥terminal.go的使用需要websocket，可使用新版本的postman，如需使用webssh 请配置正确的账户（apis/terminal.go中） 支持password 或publickey
***
## 3.Sample
* 查看k8s集群的版本:`http://127.0.0.1:10010/cluster/version/default`
```json
{
  "major": "1",
  "minor": "18",
  "gitVersion": "v1.18.2",
  "gitCommit": "52c56ce7a8272c798dbc29846288d7cd9fbae032",
  "gitTreeState": "clean",
  "buildDate": "2020-04-16T11:48:36Z",
  "goVersion": "go1.13.9",
  "compiler": "gc",
  "platform": "linux/amd64"
}
```
***
* 查看k8s集群节点详情 `http://127.0.0.1:10010/cluster/nodes/default`
```json
[
    {
        "name": "master",
        "status": "True",
        "taints": null,
        "os_image": "CentOS Linux 7 (Core)",
        "internal_ip": "10.0.8.10",
        "kernel_version": "3.10.0-1160.45.1.el7.x86_64",
        "kubelet_version": "v1.18.2",
        "creation_timestamp": "2022-04-05T03:28:46+08:00",
        "container_runtime_version": "docker://19.3.8"
    },
    {
        "name": "node1",
        "status": "True",
        "taints": null,
        "os_image": "CentOS Linux 7 (Core)",
        "internal_ip": "10.0.8.12",
        "kernel_version": "3.10.0-1160.45.1.el7.x86_64",
        "kubelet_version": "v1.18.2",
        "creation_timestamp": "2022-04-05T03:31:54+08:00",
        "container_runtime_version": "docker://19.3.8"
    }
]
```
***
* 查看k8s集群状态 `http://127.0.0.1:10010/cluster/extra/info/default`
```json lines
{
    "used_cpu": 2.0100000000000002, //已使用的cpu
    "total_cpu": 10,                //总cpu
    "used_memory": 1648361472,      //已使用的内存
    "total_memory": 15473422336,    //总内存
    "readyNodeNum": 4,              //就绪节点数量
    "totalNodeNum": 4               //总节点数量
}
```
***



