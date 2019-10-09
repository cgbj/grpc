package main

import (
    "fmt"
    //etcd_client "github.com/coreos/etcd/clientv3"
    etcd_client "go.etcd.io/etcd/clientv3"
    "time"
)

func main() {

    cli, err := etcd_client.New(etcd_client.Config{
        Endpoints:   []string{"127.0.0.1:2379"},
        DialTimeout: 5 * time.Second,
    })
    if err != nil {
        fmt.Println("connect failed, err:", err)
        return
    }

    fmt.Println("connect succ")
    defer cli.Close()
}
