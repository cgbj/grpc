# grpc
grpc and etcd3(rpc 负载均衡)

均衡功能来源
https://www.tuicool.com/articles/7NzYzuZ

使用grpc
|- helloworld grpc中的helloworld范例
|- etcdv3 注册/获均衡的地址 go
|- go_client
|- go_server 开启服务 -port 50000  带参数设置端口号
|- vendor govendor产生的包


安装说明

安装grpc
git clone https://github.com/grpc/grpc-go.git $GOPATH/src/google.golang.org/grpc
git clone https://github.com/golang/net.git $GOPATH/src/golang.org/x/net
git clone https://github.com/golang/text.git $GOPATH/src/golang.org/x/text
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
git clone https://github.com/google/go-genproto.git $GOPATH/src/google.golang.org/genproto
cd $GOPATH/src/
go install google.golang.org/grpc


go的工作目录
helloworld工作目录下
protoc -I helloworld/ helloworld/helloworld.proto --go_out=plugins=grpc:helloworld



安装 etch
    开启服务 etcd
    后台运行 nohup etcd >/tmp/etcd.log 2>&1 &
    测试功能
        etcdctl put mykey 111
        etcdctl get mykey
        etcdctl watch mykey



使用 ETCD clientV3 做服务器发现/注册  负载均衡测试
原理
    设置一个server key, watch 这个key值, 如果有新的接口, 注册服务, 如果停止服务器注销服务, 客户获取哪一个sever时候, 随机给一个

    go get -u -v github.com/coreos/etcd/clientv3
    如果下载不了
    可以到 github.com/coreos/etcd 地址下载所有的包，然后解压缩到 src\github.com\coreos 路径下，如果没有该目录则创建，并将解压后的文件夹命名为 etcd（原来为etcd-master），再将前面改名后的 etcd文件夹拷贝到 src\go.etcd.io 目录下，再使用测试程序测试下（测试前记着启动etcd的server端，同时测试程序 import "go.etcd.io/etcd/clientv3"）。

    go run etcd3_test.go


负载均衡 
    1. 创建grpc service 时候, 写一个带前缀的etcd的 key => value, 关闭service 时候删除etcd key
    2. 客户端访问时候使用负载均衡去读取对应的host

    


如果出现以下错误
    /debug/requests is already registered. You may have two independent copies of golang.org/x/net/trace in your binary, trying to maintain separate state. This may involve a vendored copy of golang.org/x
    找到问题就很好解决了，直接百度  go依赖管理-govendor
    进入项目中，govendor init初始化一下，程序会自动生成一个vendor目录
    最重要的一步来了   govendor add +external
    使用这个会把你所需要的包全部放入刚才的vendor目录中
    这个时候你go build 运行程序就不会发生panic了!!



