package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
)

var ip string

const MIN = 1
const MAX = 65535

func main() {
	flag.StringVar(&ip, "ip address", "192.168.77.3", "-ip")
	g := sync.WaitGroup{}
	for i := MIN; i < MAX; i++ {
		go func(i int) {
			g.Add(1)
			defer g.Done()
			con, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", ip, i), time.Second)
			fmt.Println(err)
			if err == nil {
				fmt.Sprintf("TCP %d is open\n", i)
				con.Close()
			}
			// con2, err := net.Dial("udp", fmt.Sprintf("%s:%d", ip, i))
			// if err == nil {
			// 	fmt.Sprintf("UDP %d is open\n", i)
			// 	con2.Close()
			// }
		}(i)
	}
	g.Wait()
}
