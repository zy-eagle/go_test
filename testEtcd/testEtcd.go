package main

import (
	"context"
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.231.129:2379", "192.168.231.129:22379", "192.168.231.129:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	_, err = cli.Put(ctx, "foo", "vvv")
	if err != nil {
		fmt.Printf("get from etcd failed,err %v\n", err)
	}

	for i := 0; i < 5; i++ {
		_, err := cli.Put(ctx, fmt.Sprintf("foo%d", i), fmt.Sprintf("bar%d", i))
		if err != nil {
			fmt.Printf("put failed, err:%v\n", err)
			return
		}
	}

	resp, err := cli.Get(ctx, "foo", clientv3.WithPrefix())
	if err != nil {
		fmt.Printf("get from etcd failed,err %v\n", err)
	}

	for _, kv := range resp.Kvs {
		fmt.Printf("%s:%s:%d \n", kv.Key, kv.Value, kv.Version)
	}
}
