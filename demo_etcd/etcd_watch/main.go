package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

// etcd watch

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error！
		fmt.Print("connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()
	// watch
	// 派一个哨兵 一直监视着 Quin 这个key的变化(新增、修改、删除)
	ch := cli.Watch(context.Background(), "Quin")
	// 从通道尝试取值(监视的信息)
	for wresp := range ch {
		for _, evt := range wresp.Events {
			fmt.Printf("Type:%v key:%s value:%s\n", evt.Type, evt.Kv.Key, evt.Kv.Value)
		}
	}

}
