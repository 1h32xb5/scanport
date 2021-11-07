package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
)

func worker(ip string,ports chan int,wg *sync.WaitGroup){
	for p := range ports{ 				

		address := fmt.Sprintf("%s:%d",ip, p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			//fmt.Printf("%s端口关闭了\n", address)
			return
		}
		conn.Close()
		fmt.Printf("%s端口打开了\n", address)

		wg.Done()

	}
}
func main(){

	var ip string
	flag.StringVar(&ip, "ip", "", "输入你要扫描端口的ip")
	flag.Parse()


	ports := make(chan int,65535)     

	var wg sync.WaitGroup

	for i := 0;i<cap(ports);i++{
		go worker(ip,ports,&wg)
	}

	for i := 1;i<65535;i++{					
		wg.Add(1)
		ports <- i      		
	}
	wg.Wait()
	close(ports)
}
