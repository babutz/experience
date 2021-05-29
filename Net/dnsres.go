package main

import (
	"fmt"
	"net"
)

func main() {
	ips, err := net.LookupIP("yandex.ru")
	if err != nil {
		panic(err)
	}

	for _, ip := range ips {
		fmt.Printf("DNS result: IN A %s\n", ip.String())
	}
}
