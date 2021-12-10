package main

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

func getConnect(zkList []string) (conn *zk.Conn) {
	conn, _, err := zk.Connect(zkList, 10*time.Second)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func main() {
	zkList := []string{"localhost:2181"}
	conn := getConnect(zkList)

	defer conn.Close()

	var flags int32 = 0
	conn.Create("go_servers", nil, flags, zk.WorldACL(zk.PermAll))

	time.Sleep(20 * time.Second)

}
