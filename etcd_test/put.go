package main

import (
    "context"
    "flag"
    "fmt"
    //etcd_client "github.com/coreos/etcd/clientv3"
    etcd_client "go.etcd.io/etcd/clientv3"
    "time"
)

var (
    keyName = flag.String("k", "aa", "set key name")
)

func main() {

    flag.Parse()

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

    flag.Parse()

    ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
    resp, err := cli.Put(ctx, "aa", "sample_value")
    cancel()
    if err != nil {
        fmt.Println("put data failed, err:", err)
        return
    }

    fmt.Println(resp)





}
