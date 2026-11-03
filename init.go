package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	local_host  string
	local_port  string
	remote_host string
	remote_port string
)

func init() {
	flag.StringVar(&local_host, "lh", "", "local host")
	flag.StringVar(&local_port, "lp", "", "local port")
	flag.StringVar(&remote_host, "rh", "", "remote host")
	flag.StringVar(&remote_port, "rp", "", "remote port")

	flag.Parse()

	if local_port == "" || remote_port == "" {
		fmt.Println("local port and remote port must be specified")
		os.Exit(1)
	}
}
