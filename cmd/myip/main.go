package main

import (
	"flag"
	"fmt"

	"github.com/hidetzu/myip"
)

var (
    remoteAddr *bool = flag.Bool("remote", false, "Show remote addrs")
)

func main() {
    flag.Parse()

    if *remoteAddr {
        remote := myip.GetRemoteAddr()
        fmt.Printf("%s", remote)
    } else {
        ip, err := myip.GetLocalIP()
        if err != nil {
            fmt.Println(err)
        }
        fmt.Println(ip)
    }
}
