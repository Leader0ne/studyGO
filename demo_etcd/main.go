package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

// etcd client put/get demo
// use etcd/clientv3
func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error！
		fmt.Printf("connect to etcd failed, err:%v\n",err)
		return
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()
	// put
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	_, err = cli.Put(ctx, "ywwuyi", "chunan")
	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed,err:%v\n", err)
		return
	}
	// get
	ctx, cancel = context.WithTimeout(context.Background(),2*time.Second)
	resp, err := cli.Get(ctx,"ywwuyi")
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed,err:%v\n", err)
		return
	}
	for _,ev := range resp.Kvs{
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	}
}
